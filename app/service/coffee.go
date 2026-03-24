package service

import (
	"errors"
	"smart-coffee/domain"
	"smart-coffee/repository"
)

var ErrNotFound = errors.New("coffee not found")

type CoffeeRepository interface {
	FindByID(id string) (domain.Coffee, error)
	Upsert(coffee domain.Coffee) error
}

type CoffeeService struct {
	repo CoffeeRepository
}

func NewCoffeeService(repo CoffeeRepository) *CoffeeService {
	return &CoffeeService{repo: repo}
}

func (s *CoffeeService) GetCoffee(id string) (domain.Coffee, error) {
	coffee, err := s.repo.FindByID(id)
	if errors.Is(err, repository.ErrNotFound) {
		return domain.Coffee{}, ErrNotFound
	}
	return coffee, err
}

func (s *CoffeeService) PutCoffee(coffee domain.Coffee) error {
	return s.repo.Upsert(coffee)
}
