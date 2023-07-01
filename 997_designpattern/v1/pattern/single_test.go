package pattern

import (
	"fmt"
	"sync"
	"testing"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance(wg *sync.WaitGroup) *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Normal : Object have been created")
			singleInstance = &single{}
		} else {
			fmt.Println("Normal : Object have been already created")
		}
	} else {
		fmt.Println("Normal : Object have been already created")
	}
	wg.Done()
	return singleInstance
}

var once sync.Once

func getInstanceWithOnce(wg *sync.WaitGroup) *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Once : Creating single instance now")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Once : Single instance already created.")
	}
	wg.Done()
	return singleInstance
}

func TestSingleton(t *testing.T) {

	ws := sync.WaitGroup{}

	for i := 0; i < 30; i++ {
		ws.Add(1)
		go getInstance(&ws)
	}

	for i := 0; i < 30; i++ {
		ws.Add(1)
		go getInstanceWithOnce(&ws)
	}

	ws.Wait()
}
