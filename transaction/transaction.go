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

package transaction

import (
	"context"
	"database/sql"
	"runtime/debug"

	"github.com/openshift-online/ocm-sdk-go/logging"
)

// By default do no roll back transaction only if it was set explicitly using rollback.Flag(ctx).
const defaultRollbackPolicy = false

// ManagerBuilder contains the data and logic needed to create
// a transaction manager.
type ManagerBuilder struct {
	logger logging.Logger
	db     *sql.DB
}

// Manager knows how to start, resolve and in general manage database
// transactions.
type Manager struct {
	logger logging.Logger
	db     *sql.DB
}

// Unit represents an sql transaction unit
type Unit struct {
	// Reference to the database transaction:
	tx *sql.Tx

	// Flag indicating if the transaction should be rolled back:
	rollbackFlag bool

	// List of functions to call after committing the transaction:
	postCommitCallbacks []func()

	// List of functions to call after rolling back the transaction:
	postRollbackCallbacks []func()
}

// Resolve resolves the current transaction according to the rollback flag.
func (u *Unit) Resolve() error {
	if u.rollbackFlag {
		if err := u.rollback(); err != nil {
			return err
		}
	} else {
		if err := u.commit(); err != nil {
			return err
		}
	}
	return nil
}

// commit commits the transaction stored in context or returns an err if one occurred.
func (u *Unit) commit() error {
	// Commit the transaction:
	err := u.tx.Commit()
	if err != nil {
		return err
	}

	// Run the post commit callbacks:
	for _, callback := range u.postCommitCallbacks {
		if callback != nil {
			go callback()
		}
	}

	return nil
}

// rollback rollbacks the transaction stored in context or returns an err if one occurred..
func (u *Unit) rollback() error {
	// Rollback the transaction:
	err := u.tx.Rollback()
	if err != nil {
		return err
	}

	// Run the post rollback callbacks:
	for _, callback := range u.postRollbackCallbacks {
		if callback != nil {
			go callback()
		}
	}

	return nil
}

// MarkForRollback flags the transaction stored in the context for rollback.
func (u *Unit) MarkForRollback() {
	u.rollbackFlag = true
}

// AddPostCommitCallback adds a callback function that will be executed after the transaction is
// successfully committed. If multiple callbacks functions are added then the order of their
// execution isn't guaranteed, and they may run in parallel in different goroutines.
func (u *Unit) AddPostCommitCallback(callback func()) error {
	u.postCommitCallbacks = append(u.postCommitCallbacks, callback)
	return nil
}

// AddPostRollbackCallback adds a callback function that will be executed after the transaction is
// rolled back. If multiple callbacks functions are added then the order of their
// execution isn't guaranteed, and they may run in parallel in different goroutines.
func (u *Unit) AddPostRollbackCallback(callback func()) error {
	u.postRollbackCallbacks = append(u.postRollbackCallbacks, callback)
	return nil
}

// NewManager creates a builder that can then be used to configure and
// create a new transaction manager.
func NewManager() *ManagerBuilder {
	return &ManagerBuilder{}
}

// Logger sets the object that the manager will use to write messages to the log.
func (b *ManagerBuilder) Logger(value logging.Logger) *ManagerBuilder {
	b.logger = value
	return b
}

// Handle sets the database handle that the manager will use to create transactions.
func (b *ManagerBuilder) Handle(value *sql.DB) *ManagerBuilder {
	b.db = value
	return b
}

// Build uses the data stored in the builder to create a new transaction manager.
func (b *ManagerBuilder) Build() (result *Manager, err error) {
	return &Manager{
		logger: b.logger,
		db:     b.db,
	}, nil
}

func (m *Manager) Begin() (*Unit, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	return &Unit{
		tx:           tx,
		rollbackFlag: defaultRollbackPolicy,
	}, nil
}

// RecoverPanic recovers from a panic if one such occurred.
func (m *Manager) RecoverPanic(ctx context.Context, failure interface{}) {
	if failure != nil {
		debug.PrintStack()
		u, err := FromContext(ctx)
		if err != nil {
			return
		}
		// Rollback transaction to avoid committing partial writes.
		u.MarkForRollback()
		err = u.Resolve()
		if err != nil {
			if err == sql.ErrTxDone {
				return
			}
			return
		}
		panic(failure)
	}
	return
}

// Close closes the connection to the DB.
func (m *Manager) Close() error {
	return m.db.Close()
}

func (m *Manager) CheckConnection() error {
	u, err := m.Begin()
	if err != nil {
		return err
	}
	ctx, err := IntoContext(u, context.Background())
	if err != nil {
		return err
	}
	defer u.Resolve()
	rows, err := m.db.QueryContext(ctx, "SELECT 1")
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
