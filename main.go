package main

import (
	"dudeStore/dudeStore/config"
	"dudeStore/dudeStore/data"
	"fmt"
)

func main() {
	koneksi := config.InitSQL()
	mdl := data.Model{}
	mdl.SetSQLConnection(koneksi)

	var username string
	var password string

	fmt.Println("Silahkan login")
	fmt.Println("Masukan username")
	fmt.Scanln(&username)
	fmt.Println("Masukan password")
	fmt.Scanln(&password)

	if username == "admin" {
		var menu int
		for menu != 99 {
			fmt.Println("Selamat datang admin")
			fmt.Println("1.Tambahkan Pegawai")
			fmt.Println("2.Hapus Pegawai")
			fmt.Println("3.Daftar Pegawai")
			fmt.Println("4.Tambahkan Produk")
			fmt.Println("5.update Produk")
			fmt.Println("6.Hapus Produk")
			fmt.Println("7.Tambahkan Pelanggan")
			fmt.Println("8.Hapus Pelanggan")
			fmt.Println("9.Tambahkan Transaksi")
			fmt.Println("10.Daftar Transaksi")
			fmt.Println("11.Hapus Transaksi")
			fmt.Println("0.Log Out")
			fmt.Scanln(&menu)
			if menu == 0 {
				break
			}
		}
	} else {
		var menu int
		for menu != 99 {
			fmt.Println("Selamat datang pegawai")
			fmt.Println("1.Tambahkan Pelanggan")
			fmt.Println("2.Tambahkan Produk")
			fmt.Println("3.Update Produk")
			fmt.Println("4.Daftar Produk")
			fmt.Println("5.Daftar Pelanggan")
			fmt.Println("0.Log Out")
			fmt.Scanln(&menu)
			if menu == 0 {
				break
			}
		}
	}
}
