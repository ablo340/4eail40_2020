package main

import (
	//"fmt"
	//"bytes"
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (s spaceEraser) Read(b []byte) (int, error){
	n, err := s.r.Read(b)

	slice := make ([]byte, n)
	copy(slice, b[0:n])
	d := strings.ReplaceAll(string(slice), " ", "")

	n = copy(b, []byte(d))

	return n, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
