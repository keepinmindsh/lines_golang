package service

import (
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	mock_currency "go_tdd/domain/currency/mockgen"
	"testing"
)

func Test_GetSumCurrencyFromInput(t *testing.T) {
	t.Run("Get Sum Currency", func(t *testing.T) {
		t.Run("Currency Mocking - CalculateStockWithPrice", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCurrencyService := mock_currency.NewMockCurrencyService(ctrl)

			mockCurrencyService.EXPECT().CalculateStockWithPrice(25, 1000.0).Return(25000.0).AnyTimes()

			price := mockCurrencyService.CalculateStockWithPrice(25, 1000.0)

			assert.Equal(t, 25000.0, price)
		})

		t.Run("Currency Mocking - GetPriceWithCurrency", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCurrencyService := mock_currency.NewMockCurrencyService(ctrl)

			// 1 USD - 1298.66 KRW
			mockCurrencyService.EXPECT().GetPriceWithCurrency("USD", 25000.0).Return(32450000.0).AnyTimes()

			price := mockCurrencyService.GetPriceWithCurrency("USD", 25000.0)

			assert.Equal(t, 32450000.0, price)
		})

		t.Run("Currency Mocking - GetSumAllPriceWithStock", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCurrencyService := mock_currency.NewMockCurrencyService(ctrl)

			mockCurrencyService.EXPECT().GetSumAllPriceWithStock([]float64{32450000.0, 32450000.0}).Return(64900000.0).AnyTimes()

			price := mockCurrencyService.GetSumAllPriceWithStock([]float64{32450000.0, 32450000.0}...)

			assert.Equal(t, 64900000.0, price)
		})

		// 각 단위 메소드를 빠르게 테스트할 수 있도록 실제 코드에 대한 구현을 정의 한다.
		t.Run("Unit Test - CalculateStockWithPrice", func(t *testing.T) {
			currencySvc := NewCurrencyService()

			stockWithPrice := currencySvc.CalculateStockWithPrice(250, 25000.0)

			assert.Equal(t, 25000.0, stockWithPrice)
		})

		t.Run("Unit Test - GetPriceWithCurrency", func(t *testing.T) {
			currencySvc := NewCurrencyService()

			currencyPrice := currencySvc.GetPriceWithCurrency("USD", 25000.0)

			assert.Equal(t, 32450000.0, currencyPrice)
		})

		t.Run("Unit Test - GetSumAllPriceWithStock", func(t *testing.T) {
			currencySvc := NewCurrencyService()

			sumPrice := currencySvc.GetSumAllPriceWithStock([]float64{32450000.0, 32450000.0}...)

			assert.Equal(t, 64900000.0, sumPrice)
		})

		t.Run("Integration Test - GetSumCurrency", func(t *testing.T) {
			currencySvc := NewCurrencyService()

			stockWithPrice := currencySvc.CalculateStockWithPrice(250, 25000.0)

			currencyPrice := currencySvc.GetPriceWithCurrency("USD", stockWithPrice)

			sumPrice := currencySvc.GetSumAllPriceWithStock([]float64{currencyPrice, currencyPrice}...)

			assert.Equal(t, 64900000.0, sumPrice)
		})
	})

}
