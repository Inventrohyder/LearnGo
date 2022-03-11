package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
	readLength, err := rot.r.Read(p)
	if err != nil {
		return 0, err
	}
	count := 0
	for i, val := range p[:readLength] {
		var reference uint8 = 0
		switch {
		case val >= 'a' && val <= 'z':
			reference = 'a'
		case val >= 'A' && val <= 'Z':
			reference = 'A'
		}

		// Only apply the rotation to A-Z, and a-z
		if reference != 0 {
			p[i] = reference + ((p[i]-reference)+13)%26
		}
		count++
	}
	return count, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
