package main

import (
	"context"
	"fmt"
)

type (
	DeployActivityParam struct {
		Repo *Repository
	}
)

func Deploy(ctx context.Context, params DeployActivityParam) error {
	fmt.Println(params.Repo.Repository.Get())
	return nil
}
