package data

import (
	"context"

	"github.com/openshift-online/ocm-sdk-go/logging"
)

var (
	ctx = context.Background()
	log = logging.Logger{}
)

func ok() {
	log.Info(ctx, "hello %s %d", "world", 42) // ok
	log.Warn(ctx, "pi is %.2f", 3.14)         // ok (placeholder with precision)
	log.Fatal(ctx, `%% %d`, 1)                // ok, percent symbol and one parameter
	log.Fatal(ctx, `%d %% %d`, 1, 2)          // ok, percent symbol between two parameters
	log.Fatal(ctx, `%d %d %%`, 1, 2)          // ok, percent symbol as last character
}

func mismatch() {
	log.Error(ctx, "name=%s age=%d", "alice")    // want "number of format placeholders .* does not match"
	log.Error(ctx, "name=%s %% age=%d", "alice") // want "number of format placeholders .* does not match"
}

func dynamicFormatIgnored() {
	f := "x=%d"
	log.Info(ctx, f, 1) // format is not a literal -> no errors expected
}

func nolintSameLine() {
	// linting is disabled: no error expected although the parameters are wrong
	log.Info(ctx, "u=%d v=%d", 1) //nolint:ocmlogger
}

func nolintPrevLine() {
	// linting is disabled: no error expected although the parameters are wrong
	//nolint:ocmlogger
	log.Info(ctx, "x=%d y=%d", 1)
}

func nolintDisabledPreviousButNotCurrent() {
	// linting is disabled: no error expected although the parameters are wrong
	//nolint:ocmlogger
	log.Info(ctx, "x=%d y=%d", 1)
	log.Info(ctx, "x=%d y=%d", 1) //nolint:ocmlogger
	log.Info(ctx, "x=%d y=%d", 1) // want "number of format placeholders .* does not match"

}
