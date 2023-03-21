package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitSQL() *sql.DB {

	db, err := sql.Open("mysql", "root:lazuardi12@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if db.Ping() != nil {
		fmt.Println(db.Ping().Error())
		return nil
	}
	defer db.Close()
	fmt.Println("yes")
	return db
}
