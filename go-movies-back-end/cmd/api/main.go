package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	//set application config
	var app application

	//read from command line

	//connect to the database

	app.Domain = "varunupadhya.com"
	log.Println("Starting application on port", port)

	//We configured the third party mux called chi and we are not using the default mux
	//http.HandleFunc("/", Hello)
	//start a web server

	//If we pass nil to the listen and serve, it will
	//create a Default MUX
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
