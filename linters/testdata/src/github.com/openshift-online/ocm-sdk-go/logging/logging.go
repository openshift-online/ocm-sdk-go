package logging

import "context"

/* logger stub */

type Logger struct{}

func (Logger) Debug(context.Context, string, ...any) {}
func (Logger) Info(context.Context, string, ...any)  {}
func (Logger) Warn(context.Context, string, ...any)  {}
func (Logger) Error(context.Context, string, ...any) {}
func (Logger) Fatal(context.Context, string, ...any) {}
