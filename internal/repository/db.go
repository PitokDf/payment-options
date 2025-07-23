package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB1 *sql.DB
var DB2 *sql.DB

func InitDB() error {
	var err error

	dsn1 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "", "localhost", "3306", "produk")

	DB1, err = sql.Open("mysql", dsn1)

	if err != nil {
		return err
	}

	dsn2 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "", "localhost", "3306", "payment")

	// open the second database connection
	DB2, err = sql.Open("mysql", dsn2)

	// check if the connection is established
	if err != nil {
		return err
	}

	// ping the database connections to check if they are alive
	// and to establish the connection
	if err = DB1.Ping(); err != nil {
		return err
	}
	if err = DB2.Ping(); err != nil {
		return err
	}
	return nil
}
