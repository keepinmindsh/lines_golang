package service

import "go_tdd/domain/currency"

type currencyService struct {
}

func NewCurrencyService() currency.CurrencyService {
	return &currencyService{}
}

func (c currencyService) CalculateStockWithPrice(stocks int, price float64) float64 {
	return 1
}

func (c currencyService) GetPriceWithCurrency(currency string, price float64) float64 {
	return 1
}

func (c currencyService) GetSumAllPriceWithStock(f ...float64) float64 {
	return 1
}
