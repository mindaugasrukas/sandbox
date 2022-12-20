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

func (s Service3) DoSometingWithServiceA() error {
	fmt.Println("Service3.DoSometingWithServiceA()")
	s.svc.DoSomething()
	return nil
}

func (s Service3) DoSomething() error {
	fmt.Println("Service3.DoSomething()")
	return nil
}
