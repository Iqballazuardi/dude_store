package main

import (
	"dudeStore/dudeStore/config"
	"dudeStore/dudeStore/data"
)

func main() {
	koneksi := config.InitSQL()
	mdl := data.Model{}
	mdl.SetSQLConnection(koneksi)

}
