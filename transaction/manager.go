/*
Copyright (c) 2020 Red Hat, Inc.

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
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
)

// ManagerBuilder contains the information and logic needed to build a transaction manager. Don't
// create instances of this type directly, use the NewManager function instead.
type ManagerBuilder struct {
	logger sdk.Logger
	db     *sql.DB
}

// Manager contains the information and logic to manage transactions.
type Manager struct {
	logger sdk.Logger
	db     *sql.DB
}

// NewManager creates a new transaction manager builder that can then be used to configure and
// create new transaction managers.
func NewManager() *ManagerBuilder {
	return &ManagerBuilder{}
}

// Logger sets the logger that the transaction manager will use to send messages to the log.
func (b *ManagerBuilder) Logger(value sdk.Logger) *ManagerBuilder {
	b.logger = value
	return b
}

// DB sets the database handle that the manager will use to create transactions.
func (b *ManagerBuilder) DB(value *sql.DB) *ManagerBuilder {
	b.db = value
	return b
}

// Build uses the information stored in the builder to create a new transaction manager.
func (b *ManagerBuilder) Build() (result *Manager, err error) {
	// Check parameters:
	if b.logger == nil {
		err = fmt.Errorf("logger is mandatory")
		return
	}
	if b.db == nil {
		err = fmt.Errorf("database handle is mandatory")
		return
	}

	// Create and populate the object:
	result = &Manager{
		logger: b.logger,
		db:     b.db,
	}

	return
}

// Begin creates a new transaction.
func (b *Manager) Begin(ctx context.Context) (result *Object, err error) {
	// Try to create the underlying database transaction:
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// Create and populate the transaction object:
	result = &Object{
		manager:  b,
		tx:       tx,
		rollback: defaultRollbackFlag,
	}

	return
}

// Complete completes the given transaction. If the transaction has been marked for roll back (with
// the MarkForRollback method) then it will be rolled back. Otherwise it will be committed.
func (m *Manager) Complete(ctx context.Context, object *Object) error {
	if object.rollback {
		err := m.rollback(object)
		if err != nil {
			m.logger.Error(ctx, "Can't rollback transaction: %v", err)
			return err
		}
		m.logger.Debug(ctx, "Rolled back transaction")
	} else {
		err := m.commit(object)
		if err != nil {
			m.logger.Error(ctx, "Can't commit transaction: %v", err)
			return err
		}
		m.logger.Debug(ctx, "Committed transaction")
	}
	return nil
}

// commit tries to commit the underlying database transaction, and if that succeeds it then runs all
// the post commit callbacks.
func (b *Manager) commit(object *Object) error {
	// Try to commit the underlying database transaction:
	err := object.tx.Commit()
	if err != nil {
		return err
	}

	// Run the post commit callbacks:
	for _, callback := range object.postCommitCallbacks {
		if callback != nil {
			go callback()
		}
	}

	return nil
}

// rollback tries to rollback the underlying database transaction, and if that succeeds it then runs
// all the post rollback callbacks.
func (b *Manager) rollback(object *Object) error {
	// Try to rollback the transaction:
	err := object.tx.Rollback()
	if err != nil {
		return err
	}

	// Run the post rollback callbacks:
	for _, callback := range object.postRollbackCallbacks {
		if callback != nil {
			go callback()
		}
	}

	return nil
}

// Default rollback flag for new transactions:
const defaultRollbackFlag = false
