package samples

import (
	"fmt"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

type Watchdog struct{ last int64 }

func (w *Watchdog) WrongKeepAlive() {
	w.last = time.Now().UnixNano()
}

func (w *Watchdog) WrongStart() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

func (w *Watchdog) RightKeepAlive() {
	atomic.StoreInt64(&w.last, time.Now().UnixNano())
}

func (w *Watchdog) RightStart() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

func Test_PrimitiveContact(t *testing.T) {
	w := &Watchdog{
		last: 10,
	}

	w.WrongKeepAlive()
	w.WrongStart()

	w.RightKeepAlive()
	w.RightStart()
}
