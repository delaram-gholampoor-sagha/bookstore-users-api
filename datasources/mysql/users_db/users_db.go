package users_db

import (
	"database/sql"
	"fmt"
	"log"
)

// check out this link : https://github.com/golang/go/wiki/SQLInterface

var (
	Client *sql.DB
)

// by importing this package you have called the init function
func init() {
	// over the host we are about to connect / the schema that we want to use
	// we have configured this database to use utf we are gonna place the charset = utf
	// username:password@tcp(host)/
	// this is the schema we have to use : "%s:%s@tcp(%s)/%s?charset=utf8"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "1234", "127.0.0.1", "users_db")
	var err error
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
