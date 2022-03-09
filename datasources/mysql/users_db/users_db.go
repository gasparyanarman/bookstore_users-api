package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB

	// username = os.Getenv("mysql_users_username")
	// password = os.Getenv("mysql_users_password")
	// host     = os.Getenv("mysql_users_host")
	// schema   = os.Getenv("mysql_users_schema")

	username = "root"
	password = "9a147gml"
	host     = "localhost:3306"
	schema   = "user_api"
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("successfully connected to db")
}
