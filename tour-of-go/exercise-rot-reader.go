// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by
// applying the rot13 substitution cipher to all alphabetical characters.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ExerciseRot13Reader() {

	fmt.Println("Exercise Rot13 Reader:")
	inputText := "Lbh penpxrq gur pbqr!"

	fmt.Println("input: ", inputText)

	s := strings.NewReader(inputText)
	r := rot13Reader{s}
	fmt.Print("output: ")
	io.Copy(os.Stdout, &r)
	fmt.Println()

	fmt.Println("----------------------------------------------------")

}

type rot13Reader struct {
	r io.Reader
}

func (rot13Reader *rot13Reader) Read(b []byte) (int, error) {

	nRead, err := rot13Reader.r.Read(b)

	for i := 0; i < nRead; i++ {
		c := b[i]
		switch {
		case c >= 'A' && c <= 'M' || c >= 'a' && c <= 'm':
			b[i] = c + 13
		case c >= 'N' && c <= 'Z' || c >= 'n' && c <= 'z':
			b[i] = c - 13
		}

	}

	return nRead, err
}
