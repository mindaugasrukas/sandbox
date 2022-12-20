package service2

import (
	"fmt"
	"sandbox/di/service3"
)

type (
	// Service2 is a service.
	Service2 struct {
	}
)

// NewService2 creates a new Service2.
func NewService2() *Service2 {
	return &Service2{}
}

func (s *Service2) DoSomething() error {
	fmt.Println("Service2.DoSomething()")
	return nil
}

func (s *Service2) DoSomething2(svc3 *service3.Service3) error {
	fmt.Println("Service2.DoSomething2()")
	svc3.DoSomething()
	return nil
}
