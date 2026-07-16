package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/thavel/apidep/pkg/file"
	"github.com/thavel/apidep/pkg/logger"
	"github.com/thavel/apidep/pkg/provider"
)

func main() {
	slog.SetDefault(slog.New(
		logger.NewCliHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	))

	providers := []file.Provider{
		&provider.FS{},
		&provider.Git{},
	}

	app := &cli.Command{
		Name:  "apidep",
		Usage: "API dependency manager",
		Commands: []*cli.Command{
			initCommand(),
			syncCommand(providers),
			validateCommand(),
			ciCommand(providers),
		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		slog.Error("apidep error", "err", err)
		os.Exit(1)
	}
}
