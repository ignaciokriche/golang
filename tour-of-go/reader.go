package main

import (
	"fmt"
	"io"
	"strings"
)

func ReaderDemo() {

	fmt.Println("Reader")

	r := strings.NewReader("Hello, Reader!")

	buff := make([]byte, 5)

	for {
		nRead, err := r.Read(buff)
		for i := 0; i < nRead; i++ {
			fmt.Printf("%c", buff[i])
		}

		if err != nil {
			if err != io.EOF {
				fmt.Println("error while reading:", err)
			}
			break
		}
	}

	fmt.Println()
	fmt.Println("----------------------------------------------------")

}
