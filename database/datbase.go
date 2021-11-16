package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func SetupDatabase() {
	var err error
	var DSN string = os.Getenv("DB_DSN")
	DB, err = sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	fmt.Println("connect success", DB)
	// rows,err := DB.Query("SELECT text FROM diary")
	// for rows.Next(){
	// 	var text string
	// 	rows.Scan(&text)
	// 	fmt.Println(text)
	// }
	defer DB.Close()

}
