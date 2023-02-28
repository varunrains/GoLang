package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//make a http call using the package
//Refer the GO documentation it is easy to understand

//Interfaces can be put into another interfaces
//To implement parent interface you need to
//implement all the child interfaces as well.

type logWriter struct{}

func main() {

	//To figure out how to use it
	//Please refer the documentation

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//response implements Reader interface
	//So you can easily call Read method on the body of the response.

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//The above code can be even more simplified.
	// io.Copy(Writer Interface, Reader Interface)
	//io.Copy(os.Stdout, resp.Body)

	//Because logWriter implements the Writer interface implicitly
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

// Now this method implements Writer interface because the signature matches
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote this many bytes:", len(bs))
	return len(bs), nil
}
