package main

import (
	"os"

	"github.com/urfave/cli"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	app := &cli.App{
		Name:  "oss-e2e-plane",
		Usage: "OSS E2E service",
		Action: func(ctx *cli.Context) error {
			fxOpts, err := InitFx(ctx)
			if err != nil {
				return err
			}

			var logger *zap.Logger
			var worker worker.Worker

			if err := fx.New(fxOpts, fx.Populate(&logger, &worker)).Err(); err != nil {
				return err
			}
			return RunWorker(logger, ctx, worker)
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
