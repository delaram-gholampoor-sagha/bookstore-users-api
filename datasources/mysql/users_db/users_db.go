package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// we are using this import for the open collection method

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// check out this link : https://github.com/golang/go/wiki/SQLInterface

var (
	Client *sql.DB
)

// by importing this package you have called the init function
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	var (
		username = os.Getenv("MYSQL_USERS_USERNAME")
		passowrd = os.Getenv("MYSQL_USERS_PASSWORD")
		host     = os.Getenv("MYSQL_USERS_HOST")
		schema   = os.Getenv("MYSQL_USERS_SCHEMA")
	)

	// over the host we are about to connect / the schema that we want to use
	// we have configured this database to use utf we are gonna place the charset = utf
	// username:password@tcp(host)/
	// this is the schema we have to use : "%s:%s@tcp(%s)/%s?charset=utf8"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, passowrd, host, schema)
	// var err error
	Client, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		// if we have any error we wont start the application
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	// if we reach this point it means that we have a valid database to connect
	log.Println("database successfully configured")

}
