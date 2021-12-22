/*
Copyright (c) 2021 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testing

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo/v2" // nolint
	. "github.com/onsi/gomega"    // nolint

	_ "github.com/jackc/pgx/v4/stdlib" // nolint
)

// DatabaseServer knows how to start a PostgreSQL database server inside a container, and how to
// create databases to be used for tests.
type DatabaseServer struct {
	// Temporary directory for database configuration:
	tmp string

	// Name of the tool used to create containers (podman or docker):
	tool string

	// Identifier of the container where the database server is running:
	container string

	// Host and port of the database server:
	host string
	port string

	// Database handle:
	handle *sql.DB

	// Number of databases created:
	count int

	// List of databases created, so we don't forget to remove them:
	dbs []*Database
}

// Database is a PostgreSQL database.
type Database struct {
	// Reference to the database server that owns this database:
	server *DatabaseServer

	// Database name, user and password:
	name     string
	user     string
	password string

	// List of database handles created, so we don't forget to close them:
	handles []*sql.DB
}

// MakeDatabaseServer creates a new database server.
func MakeDatabaseServer() *DatabaseServer {
	var err error

	// Check if podman or docker are available:
	tool, err := exec.LookPath("podman")
	if err != nil {
		tool, err = exec.LookPath("docker")
		Expect(err).ToNot(HaveOccurred(), "Can't find 'podman' or 'docker'")
	}

	// Generate a random password for the database admnistrator:
	password := uuid.NewString()

	// Create a temporary directory for the database configuration file. Note that we need to
	// explicitly change the permissions of this directory so that everybody can read and
	// execute. That is necessary because the database server needs that permission and it runs
	// with user 26 which most probably won't match our user or group.
	tmp, err := ioutil.TempDir("", "test-*.d")
	Expect(err).ToNot(HaveOccurred())
	err = os.Chmod(tmp, 0755) // #nosec G302
	Expect(err).ToNot(HaveOccurred())

	// Create the database configuration file. Like the directory this needs to be readable for
	// the user of the database server.
	confFile := filepath.Join(tmp, "db.conf")
	confText := `
		log_destination = 'stderr'
		log_statement = 'all'
		logging_collector = off
	`
	err = ioutil.WriteFile(confFile, []byte(confText), 0644)
	Expect(err).ToNot(HaveOccurred())

	// Start the database server:
	runOut := &bytes.Buffer{}
	runCmd := exec.Command(
		tool, "run",
		"--env", "POSTGRESQL_ADMIN_PASSWORD="+password,
		"--volume", tmp+":/opt/app-root/src/postgresql-cfg:Z",
		"--publish", "5432",
		"--detach",
		"docker.io/centos/postgresql-12-centos8:latest",
	) // #nosec G204
	runCmd.Stdout = runOut
	runCmd.Stderr = GinkgoWriter
	err = runCmd.Run()
	Expect(err).ToNot(HaveOccurred())
	container := strings.TrimSpace(runOut.String())

	// Find out the port number assigned to the database server:
	portOut := &bytes.Buffer{}
	portCmd := exec.Command(tool, "port", container, "5432/tcp") // #nosec G204
	portCmd.Stdout = portOut
	portCmd.Stderr = GinkgoWriter
	err = portCmd.Run()
	Expect(err).ToNot(HaveOccurred())
	portLines := strings.Split(portOut.String(), "\n")
	Expect(len(portLines)).To(BeNumerically(">=", 1))
	portLine := portLines[0]
	hostPort := strings.TrimSpace(portLine)
	host, port, err := net.SplitHostPort(hostPort)
	Expect(err).ToNot(HaveOccurred())
	if host == "0.0.0.0" {
		host = "127.0.0.1"
	}

	// Wait till the database server is responding:
	url := fmt.Sprintf(
		"postgres://postgres:%s@%s:%s/postgres?sslmode=disable",
		password, host, port,
	)
	handle, err := sql.Open("pgx", url)
	Expect(err).ToNot(HaveOccurred())
	Eventually(handle.Ping, 10, 1).ShouldNot(HaveOccurred())

	// Create and populate the object:
	return &DatabaseServer{
		tmp:       tmp,
		tool:      tool,
		container: container,
		host:      host,
		port:      port,
		handle:    handle,
	}
}

// Close stops the database server.
func (s *DatabaseServer) Close() {
	var err error

	// Delete all databases:
	for _, db := range s.dbs {
		db.Close()
	}

	// Get the logs of the database server:
	logsCmd := exec.Command(s.tool, "logs", s.container) // #nosec G204
	logsCmd.Stdout = GinkgoWriter
	logsCmd.Stderr = GinkgoWriter
	err = logsCmd.Run()
	Expect(err).ToNot(HaveOccurred())

	// Stop the database server:
	killCmd := exec.Command(s.tool, "kill", s.container) // #nosec G204
	killCmd.Stdout = GinkgoWriter
	killCmd.Stderr = GinkgoWriter
	err = killCmd.Run()
	Expect(err).ToNot(HaveOccurred())

	// Delete the temporary directories:
	err = os.RemoveAll(s.tmp)
	Expect(err).ToNot(HaveOccurred())
}

// MakeDatabase creates a new database.
func (s *DatabaseServer) MakeDatabase() *Database {
	var err error

	// Generate the database name and password:
	name := fmt.Sprintf("test_%d", s.count)
	user := fmt.Sprintf("test_%d", s.count)
	password := uuid.NewString()
	s.count++

	// Create the user:
	_, err = s.handle.Exec(fmt.Sprintf(
		"create user %s with password '%s'",
		user, password,
	))
	Expect(err).ToNot(HaveOccurred())

	// Create the database:
	_, err = s.handle.Exec(fmt.Sprintf(
		"create database %s owner %s;",
		name, user,
	))
	Expect(err).ToNot(HaveOccurred())

	// Create and populate the object:
	db := &Database{
		server:   s,
		name:     name,
		user:     user,
		password: password,
	}

	// Remember to remove it:
	s.dbs = append(s.dbs, db)

	return db
}

// MakeHandle creates a new database handle for this database.
func (d *Database) MakeHandle() *sql.DB {
	// Create the new handle:
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		d.user, d.password, d.server.host, d.server.port, d.name,
	)
	handle, err := sql.Open("pgx", url)
	Expect(err).ToNot(HaveOccurred())

	// Remember to close it:
	d.handles = append(d.handles, handle)

	return handle
}

// Close deletes the database.
func (d *Database) Close() {
	var err error

	// Close all the handles:
	for _, handle := range d.handles {
		_ = handle.Close() // #nosec G104
	}

	// Drop the database:
	_, err = d.server.handle.Exec(fmt.Sprintf(
		`drop database if exists %s`,
		d.name,
	))
	Expect(err).ToNot(HaveOccurred())

	// Drop the user:
	_, err = d.server.handle.Exec(fmt.Sprintf(
		`drop user if exists %s`,
		d.user,
	))
	Expect(err).ToNot(HaveOccurred())
}
