package pattern

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"sync"
	"testing"
)

func Test_Deligate(t *testing.T) {

	strings := []string{"task1", "task2", "task3", "task4"}

	var wg sync.WaitGroup

	for _, value := range strings {
		wg.Add(1)
		go func(value string) {
			requestURL := fmt.Sprintf("http://localhost:%s/%s", "8080", value)
			res, err := http.Get(requestURL)
			assert.NoError(t, err, "http get error has been occurred")

			t.Log(res.Body)

			all, err := io.ReadAll(res.Body)
			assert.NoError(t, err, "read file error")

			t.Log(string(all))

			assert.Equal(t, 200, res.StatusCode)

			wg.Done()
		}(value)
	}

	file, err := os.Open("/Users/lines/sources/02_linesgits/lines_golang/997_designpattern/samples/data.json")
	assert.NoError(t, err, "os file read open error")
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		fmt.Print(string(buffer[:bytesRead]))
	}

	wg.Wait()
}
