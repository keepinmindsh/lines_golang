# TDD 실제로 해보기 

## 기본적인 Project Layout 작성하기 

- Domain 정의하기 

## Mocking을 위한 기본적인 환경 구성하기 

```shell
go install github.com/golang/mock/mockgen@v1.6.0
```

> [Mock Gen](https://github.com/golang/mock)

## 기본적인 요구 사항을 파악하고 추상화하기 

- 주 * 가격을 계산할 수 있어야 한다. 
- 가격의 경우, 종목에 따라서 가격에 환율을 입력하여 계산해야 한다.
- 종목별로 계산된 값을 모두 더하여 저장해서 원하는 환율에 따라서 최종 값을 반환해야 한다. 

## 각 요구사항에 대한 인터페이스로 추상화를 정의한다. 

- 정의된 추상화 

```go
package currency

type (
	CurrencyService interface {
		CalculateStockWithPrice(stocks int, price float64) float64
		GetPriceWithCurrency(currency string, price float64) float64
		GetSumAllPriceWithStock(...float64) float64
	}
)

```

- 생성된 MockGen

```shell
$ mockgen -source=034_go_tdd/domain/currency/domain.go -destination=034_go_tdd/domain/currency/mockgen/domain.go
```

- 실제의 생성된 결과 이미지 

![Mockgen Samples](https://github.com/keepinmindsh/lines_golang/blob/main/034_go_tdd/mockgen_sample.png)

## 실제 테스트 코드 부터 시작해보기 

> BDD ( Behavior Driven Development  )

- Mocking에 의한 사전 테스트 코드를 작성하기 
  - Input, Output에 대한 예측 코드 작성하기 