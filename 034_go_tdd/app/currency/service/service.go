package service

import (
	"go_tdd/domain"
	"go_tdd/domain/currency"
)

type service struct {
	currency.CurrencyService
}

func NewService() domain.Service {
	s := &service{}
	s.Register()

	return s
}

func (s *service) Register() {
	s.CurrencyService = NewCurrencyService()
}
