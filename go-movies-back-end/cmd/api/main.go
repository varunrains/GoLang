package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     *sql.DB
}

func main() {
	//set application config
	var app application

	//read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost post=5432 user=postgres password=postgres dbname=movies sslmode disable timezone=UTC connect_timeout=5", "Postgres Connection string")
	flag.Parse()

	//connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = conn
	defer app.DB.Close()

	app.Domain = "varunupadhya.com"
	log.Println("Starting application on port", port)

	//We configured the third party mux called chi and we are not using the default mux
	//http.HandleFunc("/", Hello)
	//start a web server

	//If we pass nil to the listen and serve, it will
	//create a Default MUX
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
