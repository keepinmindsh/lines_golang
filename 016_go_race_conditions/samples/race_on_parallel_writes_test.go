package samples

import (
	"os"
	"testing"
)

func Test_ErrParallelWrite(t *testing.T) {
	ErrParallelWrite([]byte("Test"))
}

func Test_RightParallelWrite(t *testing.T) {
	RightParallelWrite([]byte("Test"))
}

func ErrParallelWrite(data []byte) chan error {
	res := make(chan error, 2)

	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main go routine
			// so the write races with the write below
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
		f2, err := os.Create("file2")
		if err != nil {
			res <- err
		} else {
			go func() {
				_, err = f2.Write(data)
				res <- err
				f2.Close()
			}()
		}
	}
	return res
}

func RightParallelWrite(data []byte) chan error {
	res := make(chan error, 2)

	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main go routine
			// so the write races with the write below
			_, err := f1.Write(data)
			res <- err
			f1.Close()
		}()
		f2, err := os.Create("file2")
		if err != nil {
			res <- err
		} else {
			go func() {
				_, err := f2.Write(data)
				res <- err
				f2.Close()
			}()
		}
	}
	return res
}
