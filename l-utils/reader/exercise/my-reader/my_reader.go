package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (m MyReader) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}
