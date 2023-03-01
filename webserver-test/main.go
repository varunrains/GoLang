package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello This is a home page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println((fmt.Sprintf("Number of bytes written: %d", n)))
}

func About(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, fmt.Sprintf("This is a about page and 2+2 is %d", addTwoValues(2, 2)))
	if err != nil {
		fmt.Println(err)
	}
}

// Error handling
func Divide(w http.ResponseWriter, r *http.Request) {
	x, y := 100.0, 0.0
	res, err := divideTwoValues(x, y)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Cannot divide by %f", y))
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("The result of %f / %f is %f", x, y, res))
}

func divideTwoValues(x, y float64) (float64, error) {
	if y <= 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return x / y, nil
}

func addTwoValues(x, y int) int {
	return x + y
}

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello Hi!!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println((fmt.Sprintf("Number of bytes written: %d", n)))
	// })
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	http.ListenAndServe(portNumber, nil)
}
