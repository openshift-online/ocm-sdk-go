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
	"fmt"
)

type contextKey int

const (
	transactionKey contextKey = iota
)

// IntoContext returns a new context with the transaction unit stored in it.
func IntoContext(u *Unit, ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, transactionKey, u), nil
}

// FromContext Retrieves the transaction from the context.
func FromContext(ctx context.Context) (*Unit, error) {
	transactionUnit, ok := ctx.Value(transactionKey).(*Unit)
	if !ok {
		return nil, fmt.Errorf("Could not retrieve transaction from context")
	}
	return transactionUnit, nil
}
