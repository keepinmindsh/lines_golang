package sample

import (
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"strconv"
	"testing"
	"time"
)

func Test_MapLo(t *testing.T) {
	strings := lop.Map([]int64{1, 2, 3, 4}, func(x int64, index int) string {
		return strconv.FormatInt(x, 10)
	})

	for _, value := range strings {
		t.Log(value)
	}
}

func Test_MapLoByFunction(t *testing.T) {
	strings := lo.Map([]func() int{Sample1, Sample2, Sample3, Sample4}, func(x func() int, index int) int {
		return x()
	})

	for _, value := range strings {
		t.Log(value)
	}
}

func Test_MapLoByFunctionParallel(t *testing.T) {
	strings := lop.Map([]func() int{Sample1, Sample2, Sample3, Sample4}, func(x func() int, index int) int {
		return x()
	})

	for _, value := range strings {
		t.Log(value)
	}
}

func Sample1() int {
	time.Sleep(1 * time.Second)
	return 1
}

func Sample2() int {
	time.Sleep(1 * time.Second)
	return 2
}

func Sample3() int {
	time.Sleep(1 * time.Second)
	return 3
}

func Sample4() int {
	time.Sleep(1 * time.Second)
	return 4
}
