package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// we are using this import for the open collection method

	_ "github.com/go-sql-driver/mysql"
)

// check out this link : https://github.com/golang/go/wiki/SQLInterface

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

// by importing this package you have called the init function
func init() {

	// over the host we are about to connect / the schema that we want to use
	// we have configured this database to use utf we are gonna place the charset = utf
	// username:password@tcp(host)/
	// this is the schema we have to use : "%s:%s@tcp(%s)/%s?charset=utf8"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	// var err error
	var err error
	Client, err = sql.Open("mysql", dataSourceName)

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
