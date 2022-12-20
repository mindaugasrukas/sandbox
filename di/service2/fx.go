package service2

import (
	"sandbox/di/service3"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewService2,
	fx.Annotate(
		NewService2,
		fx.As(new(service3.ServiceA)),
	),
)
