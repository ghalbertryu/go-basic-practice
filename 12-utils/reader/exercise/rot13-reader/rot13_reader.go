package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return
	}

	for i := 0; i < len(p); i++ {
		p[i] = rot13(p[i])
	}

	fmt.Printf("[ReadBufferLog] %q\n", p[:n])
	return
}

func rot13(c byte) byte {
	var ret byte
	if c >= 'A' && c <= 'Z' {
		ret = (c-'A'+13)%26 + 'A'
	} else if c >= 'a' && c <= 'z' {
		ret = (c-'a'+13)%26 + 'a'
	} else {
		ret = c
	}
	return ret
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
