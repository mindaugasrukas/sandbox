package service2

import (
	"sandbox/di/service3"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewService2,
	fx.Annotate(
		func(svc2 *Service2) *Service2 {
			return svc2
		},
		fx.As(new(Iface)),
		fx.As(new(service3.ServiceA)),
	),
)
