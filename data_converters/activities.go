package main

import (
	"context"
	"fmt"
)

type (
	DeployActivityParam struct {
		Repository IRepository
	}
)

func Deploy(ctx context.Context, params DeployActivityParam) error {
	fmt.Println(params.Repository.Get())
	return nil
}
