package main

import (
	"log"
	"path/filepath"
)

func main() {

	Example_SelectFanIn()

	ExampleDistribute()

	ExampleFanIn()

	FanoutMain()

	ExampleFibonacci()

	Example_simpleChannel()

	FindMinimumValue()

	// FileDownloadAndZipIt()
}

func FileDownloadAndZipIt() {
	urls := []string{
		"http://image.com/img01.jpg",
		"http://image.com/img02.jpg",
		"http://image.com/img03.jpg",
	}

	for _, url := range urls {
		go func(url string) {
			if _, err := Download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}

	filenames, err := filepath.Glob("*.jpg")

	if err != nil {
		log.Fatal(err)
	}

	err = WriteZip("images.zip", filenames)

	if err != nil {
		log.Fatal(err)
	}
}
