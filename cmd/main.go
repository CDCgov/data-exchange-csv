package main

// TODO: Move run() into a separate app package? That way we can organize how we define different execution envs - Thomas
// TODO: Use absolute paths from project structure for importing internal packages
import (
	"context"
	"fmt"
	"github.com/CDCgov/data-exchange-csv/cmd/internal/server"
	"io"
	"log/slog"
	"os"
	"os/signal"
)

func run(
	ctx context.Context,
	args []string,
	stdout, stderr io.Writer,
) error {
	// TODO: Add elapsed time since service starts to logging?
	log := slog.New(slog.NewTextHandler(stdout, nil))
	slog.SetDefault(log)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	log.Info("Starting DEX CSV Validation")

	server.New() // Event loop in here

	log.Info("Exiting DEX CSV Validation")

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
