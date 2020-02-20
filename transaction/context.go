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
	"fmt"
)

// FromContext returns the current transaction from the context. If there is no transaction in
// context then an error will be returned.
func FromContext(ctx context.Context) (result *Object, err error) {
	result, ok := ctx.Value(contextKey).(*Object)
	if !ok {
		err = fmt.Errorf("context doesn't contain a transaction")
		return
	}
	return
}

// ToContext creates a new context that contains the given transaction.
func ToContext(ctx context.Context, object *Object) context.Context {
	return context.WithValue(ctx, contextKey, object)
}

// contextKeyType is the type of the key used to store the current transaction in the context.
type contextKeyType int

// contextKey is the key used to store the current transaction in the context.
const contextKey contextKeyType = iota
