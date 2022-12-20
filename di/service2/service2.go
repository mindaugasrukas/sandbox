package service2

import (
	"fmt"
	"math/rand"
	"sandbox/di/service3"
)

type (
	// Service2 is a service.
	Service2 struct {
		v int
	}

	Iface interface {
		DoSomething() error
		DoSomething2(svc3 *service3.Service3) error
	}
)

// NewService2 creates a new Service2.
func NewService2() *Service2 {
	fmt.Println("Service2.NewService2()")
	return &Service2{
		v: rand.Int(),
	}
}

func (s *Service2) DoSomething() error {
	fmt.Printf("Service2.DoSomething() %p, %d\n", s, s.v)
	return nil
}

func (s *Service2) DoSomething2(svc3 *service3.Service3) error {
	fmt.Printf("Service2.DoSomething2() %p, %d\n", s, s.v)
	svc3.DoSomething()
	return nil
}
