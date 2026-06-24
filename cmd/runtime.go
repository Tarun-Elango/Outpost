package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"devbox-cli/service"
)

const defaultCommandTimeout = 5 * time.Minute

// CommandContext returns a context bounded to the lifetime of a CLI command.
// timer for the whole command
func CommandContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultCommandTimeout)
}

// shared helper for cmds to open the runtime, and call the service functions using rt.function()

// mustOpenRuntime opens the runtime and panics if it fails
func mustOpenRuntime() *service.Runtime {
	ctx, cancel := CommandContext()
	rt, err := service.Open(ctx, cancel)
	if err != nil {
		cancel() // cancel the context if the runtime fails to open
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return rt
}
