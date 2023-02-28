package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	//fmt.Println(os.Args)

	file, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, file)
	//fmt.Println(*file)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote this many bytes:", len(bs))
	return len(bs), nil
}
