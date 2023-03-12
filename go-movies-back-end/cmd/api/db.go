package main

//The underscore in from of the package import means
//Eventhough there is no direct usage within the code
//The dependent driver packages will make use of these packages
//Hence we need this underscore in front of it
import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.DSN)

	if err != nil {
		return nil, err
	}

	log.Println("Connected to postgres!")
	return connection, nil
}
