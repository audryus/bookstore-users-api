package user_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//IMporting
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/logger"
)

var (
	//Client to the Postgres
	Client *sql.DB

	username = os.Getenv("postgres_user_username")
	password = os.Getenv("postgres_user_password")
	host     = os.Getenv("postgres_user_host")
	db       = os.Getenv("postgres_user_db")
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, host, db,
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger(logger.Get())
	log.Println("Databae sucessfully configured")
}
