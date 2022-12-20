package service3

import (
	"fmt"
)

type (
	// Service3 is a service.
	Service3 struct {
		svc ServiceA
	}

	ServiceA interface {
		DoSomething() error
	}
)

// NewService3 creates a new Service3.
func NewService3(svc ServiceA) *Service3 {
	return &Service3{
		svc: svc,
	}
}

func (s *Service3) DoSometingWithServiceA() error {
	fmt.Printf("Service3.DoSometingWithServiceA() %p\n", s)
	s.svc.DoSomething()
	return nil
}

func (s *Service3) DoSomething() error {
	fmt.Printf("Service3.DoSomething() %p\n", s)
	return nil
}
