package currency

type (
	CurrencyService interface {
		CalculateStockWithPrice(stocks int, price float64) float64
		GetPriceWithCurrency(currency string, price float64) float64
		GetSumAllPriceWithStock(...float64) float64
	}
)
