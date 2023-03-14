package main

import (
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(Register),
)

func Register(w worker.Worker) {
	w.RegisterWorkflowWithOptions(MainWorkflow, workflow.RegisterOptions{Name: "Main"})
	w.RegisterWorkflowWithOptions(ChildWorkflow, workflow.RegisterOptions{Name: "Child"})
	w.RegisterActivity(Deploy)
}
