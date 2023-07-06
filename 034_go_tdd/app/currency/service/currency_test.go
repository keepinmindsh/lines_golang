package service

import (
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	mock_currency "go_tdd/domain/currency/mockgen"
	"testing"
)

func Test_GetSumCurrencyFromInput(t *testing.T) {
	t.Run("Wrong Calculate Value return", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		// Assert that Bar() is invoked.
		defer ctrl.Finish()

		mockCurrencyService := mock_currency.NewMockCurrencyService(ctrl)

		mockCurrencyService.EXPECT().CalculateStockWithPrice(25, 1000.0).Return(25000.0).AnyTimes()

		price := mockCurrencyService.CalculateStockWithPrice(25, 1000.0)

		assert.IsEqual(25000.0, price)
	})

}
