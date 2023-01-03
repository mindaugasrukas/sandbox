package main

import (
	"sandbox/di/service1"
	"sandbox/di/service2"
	"sandbox/di/service3"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		service1.Module,
		service2.Module,
		service3.Module,
		fx.Provide(
			fx.Annotate(
				func(svc2 *service2.Service2) *service2.Service2 {
					return svc2
				},
				fx.As(new(service3.ServiceA)),
			),
		),
		fx.Invoke(service1.ServiceLifecycleHooks),
	)

	app.Run()
}
