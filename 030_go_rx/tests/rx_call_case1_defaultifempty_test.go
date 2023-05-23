package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_DefaultIfEmpty(t *testing.T) {
	observable := rxgo.Empty().DefaultIfEmpty(1)

	for item := range observable.Observe() {
		t.Log(item.V)
	}

}
