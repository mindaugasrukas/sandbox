package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/oklog/run"
	"github.com/urfave/cli"
)

const (
	DebugFlag         = "debug"
	NumDefaultPollers = 4

	// workerStopTimeout is the worker graceful stop timeout
	workerStopTimeout = 30 * time.Second
)

func newWorker(log *zap.Logger) (worker.Worker, error) {
	serviceClient, err := client.NewLazyClient(client.Options{
		Namespace: "default",
		HostPort:  "localhost:7233",
	})
	if err != nil {
		log.Info("failed to build temporal client", zap.Error(err))
		return nil, err
	}

	workerErrorHandler := func(err error) {
		log.Info("temporal worker fatal error", zap.Error(err))
	}

	return worker.New(serviceClient, "dummy-tq", worker.Options{
		MaxConcurrentWorkflowTaskPollers:        NumDefaultPollers,
		MaxConcurrentActivityTaskPollers:        8 * NumDefaultPollers,
		MaxConcurrentWorkflowTaskExecutionSize:  256,
		MaxConcurrentLocalActivityExecutionSize: 256,
		MaxConcurrentActivityExecutionSize:      256,
		WorkerStopTimeout:                       workerStopTimeout,
		OnFatalError:                            workerErrorHandler,
	}), nil
}

func InitFx(ctx *cli.Context) (fx.Option, error) {
	logger, _ := zap.NewProduction()

	var workerOpts = []fx.Option{
		fx.Supply(logger),

		Module,
		// Temporal Worker
		fx.Provide(newWorker),
	}

	return fx.Options(workerOpts...), nil
}

func RunWorker(log *zap.Logger, ctx *cli.Context, w worker.Worker) error {
	// Run group is used to ensure that all essential
	// tasks running in this worker run together
	// and shutdown together
	g := &run.Group{}

	if err := addInterruptHandler(log, g); err != nil {
		return err
	}

	if err := addTemporalWorker(log, g, w); err != nil {
		return err
	}

	// Run all actors (functions) concurrently.
	// When the first actor returns, all others are interrupted.
	// Run only returns when all actors have exited.
	// Run returns the error returned by the first exiting actor.
	return g.Run()
}

func addInterruptHandler(log *zap.Logger, g *run.Group) error {
	stopCh := make(chan struct{})
	g.Add(func() error {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		select {
		case s := <-sigCh:
			log.Info("termination signal received", zap.String("Signal", s.String()))

			go func() {
				// Exit the process immediately if another interrupt signal is received
				<-sigCh
				os.Exit(1)
			}()
		case <-stopCh:
		}
		return nil
	}, func(error) {
		close(stopCh)
	})

	return nil
}

func addTemporalWorker(log *zap.Logger, g *run.Group, w worker.Worker) error {
	stopCh := make(chan struct{})
	g.Add(func() error {
		// Start non-blocking in order to control lifecycle with run group and log each phase
		log.Info("temporal worker starting")
		if err := w.Start(); err != nil {
			return err
		}
		log.Info("temporal worker started")

		// Block until interrupted
		<-stopCh
		log.Info("temporal worker stopping")

		// Blocking call to drain work
		w.Stop()
		log.Info("temporal worker stopped")

		return nil
	}, func(error) {
		close(stopCh)
	})

	return nil
}
