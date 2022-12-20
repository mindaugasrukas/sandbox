package service2

import (
	"sandbox/di/service3"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewService2,
		fx.As(new(Iface)),
		fx.As(new(service3.ServiceA)),
	),
)
