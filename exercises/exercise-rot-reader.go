package main

import (
	_ "fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (m rot13Reader) Read(x []byte) (int, error) {
	n, e := m.r.Read(x)

	for i, val := range x {
		if val >= 'A' && val <= 'Z' {
			x[i] = ((val-'A'+13)%26 + 'A')
		} else if val >= 'a' && val <= 'z' {
			x[i] = ((val-'a'+13)%26 + 'a')
		} else {
			x[i] = val
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

/*
Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

-----------------------------
|   rot13Reader    			|
|----------------------------
| + r io.Reader            <|-- input stream
|----|----------------------|
|    ->	rot13Reader.Read() -|-> output stream
|---------------------------|

*/
