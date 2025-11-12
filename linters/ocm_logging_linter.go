package linters

import (
	"bytes"
	"go/ast"
	"go/token"
	"os"
	"strconv"
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() { register.Plugin("ocmlogger", New) }

// OcmLoggerLinter checks that calls to the OCM logger (e.g. `logger.Info`, `logger.Warn`, etc.) use format strings correctly.
//
// Specifically, it verifies that the number of format specifiers (e.g. `%v`, `%s`, `%d`, ...)
// in the log message matches the number of arguments passed after the format string.
//
// Example of a valid call:
//
//	logger.Warn(ctx, "failed to create resource %s: %v", name, err)
//
// Example of an invalid call (missing one argument):
//
//	logger.Warn(ctx, "failed to create resource %s: %v", name)
//
// The analyzer only runs on calls whose receiver is of type
// `github.com/openshift-online/ocm-sdk-go/logging.Logger` (or a pointer to it).
//
// To disable the check for a specific line, add one of the following comments:
//
//	//nolint:ocmlogger
//	// ocm-linter:ignore
//
// Example:
//
//	logger.Info(ctx, "%s %d", name) //nolint:ocmlogger
//
// Comments can be placed on the same line (or the line immediately above) the call.
type OcmLoggerLinter struct{}

func New(any) (register.LinterPlugin, error) { return &OcmLoggerLinter{}, nil }

func (l *OcmLoggerLinter) GetLoadMode() string { return register.LoadModeSyntax }

func (l *OcmLoggerLinter) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{{
		Name: "ocmlogger",
		Doc:  "find ocm logging usage errors",
		Run:  l.run,
	}}, nil
}

// fileCtx holds per-file state to avoid repeated I/O and repeated scans.
type fileCtx struct {
	pass     *analysis.Pass
	file     *ast.File
	src      []byte
	disabled map[int]bool // line -> true if the check is disabled for that line
}

// buildDisabledMap pre-computes which lines are disabled via directives.
// A directive on the same line always disables that line.
// A directive on the previous line disables the next line only if the comment is standalone.
func (fc *fileCtx) buildDisabledMap() {
	fc.disabled = make(map[int]bool)
	if fc.file == nil {
		return
	}
	for _, cg := range fc.file.Comments {
		for _, c := range cg.List {
			txt := c.Text
			if !strings.Contains(txt, "nolint:ocmlogger") && !strings.Contains(txt, "ocm-linter:ignore") {
				continue
			}
			cp := fc.pass.Fset.PositionFor(c.Pos(), true)
			// same line
			fc.disabled[cp.Line] = true
			// next line if standalone
			if isStandaloneComment(fc.src, fc.pass, c) {
				fc.disabled[cp.Line+1] = true
			}
		}
	}
}

// isDisabled returns true if a node's starting line is marked as disabled.
func (fc *fileCtx) isDisabled(node ast.Node) bool {
	if node == nil {
		return false
	}
	line := fc.pass.Fset.PositionFor(node.Pos(), true).Line
	return fc.disabled[line]
}

// isStandaloneComment returns true if the comment starts on a line with only whitespace before it.
func isStandaloneComment(src []byte, pass *analysis.Pass, c *ast.Comment) bool {
	cp := pass.Fset.PositionFor(c.Pos(), true)
	tf := pass.Fset.File(c.Pos())
	if tf == nil {
		return false
	}
	// If source is unavailable, approximate: consider standalone only if column == 1.
	if src == nil {
		return cp.Column == 1
	}
	lineStart := pass.Fset.PositionFor(tf.LineStart(cp.Line), true)
	if lineStart.Offset < 0 || cp.Offset < lineStart.Offset || cp.Offset > len(src) {
		return false
	}
	prefix := src[lineStart.Offset:cp.Offset]
	return len(bytes.TrimSpace(prefix)) == 0
}

// extractString resolves a string literal or a concatenation of string literals.
func (fc *fileCtx) extractString(expr ast.Expr) (string, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind != token.STRING {
			return "", false
		}
		s, err := strconv.Unquote(e.Value)
		return s, err == nil
	case *ast.BinaryExpr:
		if e.Op != token.ADD {
			return "", false
		}
		lhs, ok1 := fc.extractString(e.X)
		rhs, ok2 := fc.extractString(e.Y)
		if ok1 && ok2 {
			return lhs + rhs, true
		}
	}
	return "", false
}

// countPlaceholders counts single '%' placeholders and skips '%%'.
func (fc *fileCtx) countPlaceholders(s string) int {
	n := 0
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] != '%' {
			continue
		}
		// skip '%%'
		if i+1 < len(b) && b[i+1] == '%' {
			i++
			continue
		}
		n++
	}
	return n
}

func (l *OcmLoggerLinter) run(pass *analysis.Pass) (interface{}, error) {
	// Map is faster than slices.Contains for a tiny static set.
	loggerMethods := map[string]struct{}{
		"Debug": {}, "Info": {}, "Warn": {}, "Error": {}, "Fatal": {},
	}

	for _, file := range pass.Files {
		filename := pass.Fset.Position(file.Package).Filename
		src, _ := os.ReadFile(filename) // best-effort; nil on failure is fine

		fc := &fileCtx{pass: pass, file: file, src: src}
		fc.buildDisabledMap()

		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			if _, ok := loggerMethods[sel.Sel.Name]; !ok {
				return true
			}

			// Check static type of the receiver.
			recvType := pass.TypesInfo.TypeOf(sel.X)
			if recvType == nil {
				return true
			}
			ts := recvType.String()
			if ts != "github.com/openshift-online/ocm-sdk-go/logging.Logger" &&
				ts != "*github.com/openshift-online/ocm-sdk-go/logging.Logger" {
				return true
			}

			// Require at least ctx and format string.
			if len(call.Args) < 2 {
				return true // will be a compile-time error anyway
			}

			// Resolve format string at analysis time.
			fmtStr, ok := fc.extractString(call.Args[1])
			if !ok {
				return true
			}

			// Compare placeholders vs variadic args after ctx and format.
			want := fc.countPlaceholders(fmtStr)
			got := len(call.Args) - 2

			if want != got && !fc.isDisabled(call) {
				pass.Reportf(call.Pos(),
					"number of format placeholders (%d) does not match number of arguments (%d)",
					want, got)
			}
			return true
		})
	}

	return nil, nil
}
