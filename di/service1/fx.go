package service1

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewService1,
)
