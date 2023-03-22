package main

import (
	"database/sql"
	"dudeStore/config"
	"dudeStore/data"
	"fmt"
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
	var id int
	var nama string
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
			res, err := mdl.Login(username, password)
			if err != nil {
				fmt.Println("password/username salah", err)
				break
			}
			fmt.Println("halo selamat datang " + res.Nama)
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
				fmt.Println("Terima kasih telah melakukan pekerjaan anda")
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
					break
				}

				fmt.Println("sukses menghapus data")
			} else if menu == 3 {
				res, err := mdl.GetAllPegawai()
				if err != nil {
					fmt.Println("Terjadi sebuah kesalahan")
					break
				}

				for i := 0; i < len(res); i++ {
					fmt.Println(res[i])
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
			}
		}
	} else {

		var menu int
		for menu != 99 {
			res, err := mdl.Login(username, password)
			if err != nil {
				fmt.Println("password/username salah", err)
				break
			}
			id = res.Id
			nama = res.Nama
			fmt.Println("halo selamat datang " + nama)
			// fmt.Println(res)
			fmt.Println("1.Tambahkan Pelanggan")
			fmt.Println("2.Tambahkan Produk")
			fmt.Println("3.Update Produk")
			fmt.Println("4.Daftar Produk")
			fmt.Println("5.Daftar Pelanggan")
			fmt.Println("6.Tambahkan Transaksi")
			fmt.Println("7.Daftar Transaksi")
			fmt.Println("0.Log Out")
			fmt.Scanln(&menu)
			if menu == 0 {
				fmt.Println("Terima kasih telah melakukan pekerjaan anda")
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
				// fmt.Print(pelanggan)
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
