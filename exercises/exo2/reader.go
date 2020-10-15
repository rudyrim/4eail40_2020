package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

func (se spaceEraser) Read(buff []byte) (int, error) { //buffer de byte
	n, err := se.r.Read(buff)
	j := 0
	for i := 0; i < n; i++ {
		if buff[i] != 32 { //space
			buff[j] = buff[i]
			j++
		}
	}
	return j, err
}
