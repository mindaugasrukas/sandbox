package main

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

type (
	ChildWorkflowParams struct {
		Repositories []Repository
	}
)

func MainWorkflow(ctx workflow.Context) error {
	childWorkflowOptions := workflow.ChildWorkflowOptions{}
	ctxChild := workflow.WithChildOptions(ctx, childWorkflowOptions)
	childWorkflowParams := ChildWorkflowParams{
		Repositories: []Repository{
			{
				Repository: &GitRepository{
					Url: "URL-1",
				},
			},
			{
				Repository: &FsRepository{
					Path: "PATH-1",
				},
			},
		},
	}

	return workflow.ExecuteChildWorkflow(ctxChild, ChildWorkflow, childWorkflowParams).Get(ctx, nil)
}

func ChildWorkflow(ctx workflow.Context, params ChildWorkflowParams) error {

	for _, r := range params.Repositories {
		fmt.Println(r.Repository.Get())
		deployActivityOptions := workflow.ActivityOptions{
			StartToCloseTimeout: time.Minute,
			HeartbeatTimeout:    0,
		}
		ctxDeploy := workflow.WithActivityOptions(ctx, deployActivityOptions)
		deployHelmChartActivityParam := DeployActivityParam{
			Repo: r,
		}

		err := workflow.ExecuteActivity(ctxDeploy, Deploy, deployHelmChartActivityParam).Get(ctxDeploy, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
