package main

import (
	"database/sql"
	"dudeStore/config"
	"dudeStore/data"
	"dudeStore/helpers"
	"fmt"
	"log"
)

func FormUpdate(mdl data.Model, db *sql.DB) {
	mdl.SetSQLConnection(db)
	var id int
	fmt.Print("Masukkan ID produk yang ingin diupdate: ")
	fmt.Scanln(&id)

	existingProduk, err := mdl.GetProdukById(id)
	if err != nil {
		fmt.Printf("Gagal mendapatkan produk: %v\n\n", err)
	} else if existingProduk == nil {
		fmt.Printf("Tidak ada produk dengan ID %d\n\n", id)
	} else {
		fmt.Printf("Produk saat ini: %+v\n", existingProduk)

		var updatedProduk data.Produk
		fmt.Print("Masukkan nama produk baru: ")
		fmt.Scanln(&updatedProduk.Nama)
		fmt.Print("Masukkan keterangan produk baru: ")
		fmt.Scanln(&updatedProduk.Keterangan)
		fmt.Print("Masukkan stok produk baru: ")
		fmt.Scanln(&updatedProduk.Stok)
		fmt.Print("Masukkan harga produk baru: ")
		fmt.Scanln(&updatedProduk.Harga)

		err = mdl.UpdateProduk(id, updatedProduk)
		if err != nil {
			fmt.Printf("Gagal mengupdate produk: %v\n\n", err)
		} else {
			fmt.Println("Produk berhasil diupdate\n")
		}
	}
}

func main() {
	koneksi := config.InitSQL()
	mdl := data.Model{}
	mdl.SetSQLConnection(koneksi)
	// ui := &LihatDaftarPegawai{mdl}
	var id int
	var nama string
	var username string
	var password string

	fmt.Println()
	helpers.StartCmd()

	fmt.Println("=========================")
	fmt.Println("   Silahkan login dude   ")
	fmt.Println("=========================")
	fmt.Printf("Masukan username : ")
	fmt.Scanln(&username)
	fmt.Printf("Masukan password : ")
	fmt.Scanln(&password)

	if username == "admin" {
		var menu int
		for menu != 99 {
			res, err := mdl.Login(username, password)
			if err != nil {
				fmt.Println("password/username salah", err)
				fmt.Println("Mohon coba lagi")
			}

			id = res.Id
			nama = res.Nama
			fmt.Println()
			fmt.Println("===========================")
			fmt.Println("Halo selamat datang " + res.Nama)
			fmt.Println("===========================")
			fmt.Println("1.Tambahkan Pegawai")
			fmt.Println("2.Hapus Pegawai")
			fmt.Println("3.Daftar Pegawai")
			fmt.Println("4.Tambahkan Produk")
			fmt.Println("5.Update Produk")
			fmt.Println("6.Daftar Produk")
			fmt.Println("7.Hapus Produk")
			fmt.Println("8.Tambahkan Pelanggan")
			fmt.Println("9.Daftar Pelanggan")
			fmt.Println("10.Hapus Pelanggan")
			fmt.Println("11.Tambahkan Transaksi")
			fmt.Println("12.Daftar Transaksi")
			fmt.Println("13.Hapus Transaksi")
			fmt.Println("0.Log Out")
			fmt.Println("=====================")
			fmt.Printf("Enter piliihan kamu : ")
			fmt.Scanln(&menu)
			fmt.Println("=====================")
			fmt.Println("")
			if menu == 0 {
				helpers.CloseCmd()
				helpers.WordBye()
				break
			} else if menu == 1 {
				err := mdl.TambahPegawai(data.Pegawai{})
				if err != nil {
					fmt.Printf("Username sudah ada, GAGAL membuat Akun\n\n")
				} else {

					fmt.Printf("Pegawai BERHASIL ditambahkan!\n\n")
				}

			} else if menu == 2 {
				var email string
				fmt.Print("Masukkan email Pegawai:")
				fmt.Scanln(&email)
				err := mdl.DeletePegawai(email)
				if err != nil {
					fmt.Println("Terjadi sebuah kesalahan")
				}

				fmt.Println("sukses menghapus data")
			} else if menu == 3 {
				err := mdl.LihatDaftarPegawai()
				if err != nil {
					fmt.Println(err)
					fmt.Println("Terjadi sebuah kesalahan")
				}
			} else if menu == 4 {
				var produk = data.Produk{}
				fmt.Print("Masukkan Nama produk: ")
				fmt.Scanln(&produk.Nama)
				fmt.Print("Masukkan keterangan produk: ")
				fmt.Scanln(&produk.Keterangan)
				fmt.Print("Masukkan Stok produk: ")
				fmt.Scanln(&produk.Stok)
				fmt.Print("Masukkan Harga produk: ")
				fmt.Scanln(&produk.Harga)
				produk.Pegawai_id = res.Id
				// fmt.Print(produk)
				err := mdl.TambahProduk(produk)
				if err != nil {
					fmt.Printf("GAGAL menahbahkan Produk\n\n")

				} else {
					fmt.Printf("Produk BERHASIL ditambahkan!\n\n")
				}
			} else if menu == 5 {
				FormUpdate(mdl, koneksi)
			} else if menu == 6 {
				err := mdl.LihatDaftarProduk()
				if err != nil {
					fmt.Println("Terjadi sebuah kesalahan")
					fmt.Println(err)
				}
			} else if menu == 7 {
				var id int
				fmt.Print("Masukkan id Produk:")
				fmt.Scanln(&id)
				err := mdl.DeleteProduk(id)
				if err != nil {
					fmt.Println("Terjadi sebuah kesalahan")
				}

				fmt.Println("sukses menghapus data")
			} else if menu == 8 {
				var pelanggan = data.Pelanggan{}
				fmt.Print("Masukkan Nama pelanggan: ")
				fmt.Scanln(&pelanggan.Nama)
				fmt.Print("Masukkan Nomor HP pelanggan: ")
				fmt.Scanln(&pelanggan.Hp)
				fmt.Print("Masukkan Alamat pelanggan: ")
				fmt.Scanln(&pelanggan.Alamat)
				pelanggan.Pegawai_id = id
				err := mdl.TambahPelanggan(pelanggan)
				if err != nil {
					fmt.Printf("GAGAL menahbahkan pelanggan\n\n")

				} else {

					fmt.Printf("pelanggan BERHASIL ditambahkan!\n\n")
				}
			} else if menu == 9 {
				err := mdl.LihatDaftarPelanggan()
				if err != nil {
					fmt.Println("Terjadi sebuah kesalahan")
					fmt.Println(err)
				}
			} else if menu == 10 {
				var id int
				fmt.Print("Masukkan id Pelanggan:")
				fmt.Scanln(&id)
				err := mdl.DeletePelanggan(id)
				if err != nil {
					fmt.Println("Terjadi sebuah kesalahan")
				}

				fmt.Println("sukses menghapus data")

			} else if menu == 11 {
				var pelanggan int
				var produk []int
				var qty []int
				err := mdl.LihatDaftarPelanggan()
				if err != nil {
					log.Fatal(err)
					continue
				}
				fmt.Println()
				fmt.Print("Masukan id pelanggan: ")
				fmt.Scanln(&pelanggan)

				err1 := mdl.LihatDaftarProduk()
				if err1 != nil {
					fmt.Println("Terjadi sebuah kesalahan")
					fmt.Println(err1)
				}
				fmt.Println()
				isrepeat := true
				var choice string
				for isrepeat {
					var produkid int
					var qtyy int
					if choice == "t" {
						break
					}
					fmt.Print("Masukan Produk: ")
					fmt.Scanln(&produkid)
					produk = append(produk, produkid)
					fmt.Print("Masukan Masukan Qty: ")
					fmt.Scanln(&qtyy)
					qty = append(qty, qtyy)
					fmt.Print("Apakah ingin menambah produk lagi? (y/t): ")
					fmt.Scanln(&choice)
				}
				prdk, _ := mdl.DaftarProduk()
				fmt.Println(produk, qty)
				for i, val := range produk {
					produks := prdk[val]
					err := mdl.InsertTransaksi(&data.Transaksi{Nama_produk: produks.Nama, Qty: qty[i], Pelanggan_id: pelanggan, Pegawai_id: id})
					if err != nil {
						fmt.Println("Error pada index", i, err)
						break
					}
				}
			}
		}
	} else {

		var menu int
		for menu != 99 {
			res, err := mdl.Login(username, password)
			if err != nil {
				fmt.Println("password/username salah", err)
			}
			id = res.Id
			nama = res.Nama
			fmt.Println()
			fmt.Println("===========================")
			fmt.Println("Halo selamat datang " + nama)
			fmt.Println("===========================")
			fmt.Println("1.Tambahkan Pelanggan")
			fmt.Println("2.Tambahkan Produk")
			fmt.Println("3.Update Produk")
			fmt.Println("4.Daftar Produk")
			fmt.Println("5.Daftar Pelanggan")
			fmt.Println("6.Tambahkan Transaksi")
			fmt.Println("7.Daftar Transaksi")
			fmt.Println("0.Log Out")
			fmt.Println("=====================")
			fmt.Printf("Enter pilihan kamu : ")
			fmt.Scanln(&menu)
			fmt.Println("=====================")
			fmt.Println()
			if menu == 0 {
				helpers.CloseCmd()
				helpers.WordBye()
				break
			} else if menu == 1 {
				var pelanggan = data.Pelanggan{}
				fmt.Print("Masukkan Nama pelanggan: ")
				fmt.Scanln(&pelanggan.Nama)
				fmt.Print("Masukkan Nomor HP pelanggan: ")
				fmt.Scanln(&pelanggan.Hp)
				fmt.Print("Masukkan Alamat pelanggan: ")
				fmt.Scanln(&pelanggan.Alamat)
				pelanggan.Pegawai_id = id
				err := mdl.TambahPelanggan(pelanggan)
				if err != nil {
					fmt.Printf("GAGAL menahbahkan pelanggan\n\n")

				} else {

					fmt.Printf("pelanggan BERHASIL ditambahkan!\n\n")
				}
			} else if menu == 2 {
				var produk = data.Produk{}
				fmt.Print("Masukkan Nama produk: ")
				fmt.Scanln(&produk.Nama)
				fmt.Print("Masukkan keterangan produk: ")
				fmt.Scanln(&produk.Keterangan)
				fmt.Print("Masukkan Stok produk: ")
				fmt.Scanln(&produk.Stok)
				fmt.Print("Masukkan Harga produk: ")
				fmt.Scanln(&produk.Harga)
				produk.Pegawai_id = id
				// fmt.Print(produk)
				err := mdl.TambahProduk(produk)
				if err != nil {
					fmt.Printf("GAGAL menahbahkan Produk\n\n")

				} else {

					fmt.Printf("Produk BERHASIL ditambahkan!\n\n")
				}
			}
		}
	}
}
