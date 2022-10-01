package main

import (
	"archive/zip"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func Download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}

	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	defer f.Close()

	_, err := io.Copy(f, resp.Body)

	return filename, err
}

func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return filepath.Base(url.Path), nil
}

func WriteZip(outFileName string, filenames []string) error {
	outf, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(outf)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}
	return zw.Close()
}
