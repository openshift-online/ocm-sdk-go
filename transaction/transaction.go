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
	"database/sql"
)

// Object contains the logic to manipulate a transaction. Note that commiting or rolling back a
// transaction can't be done via this object. To do so it is necessary to use the Complete method
// of the transaction manager that created it.
type Object struct {
	// Reference to the transaction manager that created the transaction:
	manager *Manager

	// Reference to the database transaction:
	tx *sql.Tx

	// Flag indicating if the transaction should be rolled back:
	rollback bool

	// List of functions to call after committing the transaction:
	postCommitCallbacks []func()

	// List of functions to call after rolling back the transaction:
	postRollbackCallbacks []func()
}

// Manager returns the reference to the transaction manager that created this transaction.
func (o *Object) Manager() *Manager {
	return o.manager
}

// MarkForRollback marks the transaction so that it will be rolled back.
func (o *Object) MarkForRollback() {
	o.rollback = true
}

// Tx returns the underlying database transaction.
func (o *Object) TX() *sql.Tx {
	return o.tx
}

// AddPostCommitCallback adds a callback function that will be executed after the transaction is
// successfully committed. If multiple callbacks functions are added then the order of their
// execution isn't guaranteed, and they may run in parallel in different goroutines.
func (o *Object) AddPostCommitCallback(callback func()) {
	o.postCommitCallbacks = append(o.postCommitCallbacks, callback)
}

// AddPostRollbackCallback adds a callback function that will be executed after the transaction is
// rolled back. If multiple callbacks functions are added then the order of their execution isn't
// guaranteed, and they may run in parallel in different goroutines.
func (o *Object) AddPostRollbackCallback(callback func()) {
	o.postRollbackCallbacks = append(o.postRollbackCallbacks, callback)
}
