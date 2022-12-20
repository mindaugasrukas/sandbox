package service1

import (
	"context"
	"fmt"
	"sandbox/di/service2"
	"sandbox/di/service3"

	"go.uber.org/fx"
)

type (
	// Service1 is a service.
	Service1 struct {
		svc2  *service2.Service2
		svc2i service2.Iface
		svc3  *service3.Service3
	}
)

// NewService1 creates a new Service1.
func NewService1(svc2 *service2.Service2, svc2i service2.Iface, svc3 *service3.Service3) *Service1 {
	return &Service1{
		svc2:  svc2,
		svc2i: svc2i,
		svc3:  svc3,
	}
}

func (s *Service1) Start() error {
	fmt.Printf("Service1.Start() %p\n", s)
	s.svc2.DoSomething2(s.svc3)
	return nil
}

func (s *Service1) Stop() {
	fmt.Printf("Service1.Stop() %p\n", s)
}

func ServiceLifecycleHooks(
	lc fx.Lifecycle,
	svc *Service1,
) {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				if err := svc.Start(); err != nil {
					return err
				}
				return svc.svc3.DoSometingWithServiceA()
			},
			OnStop: func(ctx context.Context) error {
				svc.Stop()
				return nil
			},
		},
	)
}
