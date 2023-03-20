package main

import (
	"dudeStore/dudeStore/config"
	"fmt"
)

func main() {
	db := config.InitSQL()
	defer db.Close()
	fmt.Println(db)
}
