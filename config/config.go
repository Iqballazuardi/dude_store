package config

import (
	"database/sql"
	"fmt"

	// "log"
	// "os"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/joho/godotenv"
)

func InitSQL() *sql.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// user := os.Getenv("user")
	// pass := os.Getenv("pass")
	// host := os.Getenv("host")
	// port := os.Getenv("port")
	// dbname := os.Getenv("dbname")

	// connectionString := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := sql.Open("mysql", connectionString)
	db, err := sql.Open("mysql", "root:lazuardi12@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if db.Ping() != nil {
		fmt.Println(db.Ping().Error())
		return nil
	}
	// defer db.Close()
	fmt.Println("yes")
	return db
}
