package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (tr spaceEraser) Read(buffer []byte) (int, error) {
	n, err := tr.r.Read(buffer)
	j := 0
	for i := 0; i < n; i++ {
		if buffer[i] != 32 {
			buffer[j] = buffer[i]
			j++
		}
	}
	return j, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}
