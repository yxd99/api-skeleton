package services

import "fmt"

type GreetService struct{}

func NewGreetService() *GreetService {
	return &GreetService{}
}

func (s *GreetService) GenerateGreet(name string) string {
	return fmt.Sprintf("Hey %s!", name)
}
