package _19_uber_style_golang_guide

import (
	"fmt"
	"sync"
	"testing"
)

// Test_CopyValueNotReferencesWithSlice
// 사용하려는 업무 로직에 따라서 값복사냐 주소값복사냐를 명확히해서 사용해야함! 아래의 경우에는 주소값 복사임.
func Test_CopyValueNotReferencesWithSlice(t *testing.T) {
	driver := &Driver{}

	driver.SampleForCallByValueWithSlices([]Trip{})
}

type Driver struct {
	trips []Trip
}

type Trip struct {
}

func (d *Driver) SampleForCallByValueWithSlices(trips []Trip) {
	d.trips = make([]Trip, len(trips))
	copy(d.trips, trips)
}

// Test_ModificationMapOrSlices
// 해당 코드는 Mutex에 의해서 보호되는 코드임. 경쟁적인 상황에서 데이터를 보호할 수 있는 방법이다.
// 만약 make르 쓰지않고 바로 s.counters를 반환하면 data race condition이 발생할 수 있음
func Test_ModificationMapOrSlices(t *testing.T) {
	stats := &Stats{}

	snapshot := stats.Snapshot()

	for s, i := range snapshot {
		fmt.Printf("%s, %d", s, i)
	}
}

type Stats struct {
	mu       sync.Mutex
	counters map[string]int
}

func (s *Stats) Snapshot() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make(map[string]int, len(s.counters))
	for k, v := range s.counters {
		result[k] = v
	}
	return result
}

// Test_DeferCleanUP
// 테스트 코드 실행 시 defer에 대한 가독성 및 정의는 release할 대상에 대해서 선언시 종료도 defer를 같이 적어주는 것이 가독성에 좋음
func Test_DeferCleanUP(t *testing.T) {
	p := &DeferCleanUp{}

	p.Lock()
	defer p.Unlock()

	if p.count < 10 {
		fmt.Println(p.count)
	}

	p.count++
	fmt.Println(p.count)
}

type DeferCleanUp struct {
	sync.Mutex
	count int
}

// Test_EnumStartAtOne
// Enum은 0이 아닌 1부터 시작할것!
func Test_EnumStartAtOne(t *testing.T) {
	fmt.Println(Add)
}

type Operation int

const (
	Add Operation = iota + 1
	Subtract
	Multiply
)
