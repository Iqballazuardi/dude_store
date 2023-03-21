package data

import "time"

type Produk struct {
	Id         int
	Nama       string
	Keterangan string
	Stok       int
	Harga      int
	Create_at  time.Time
	Update_at  time.Time
	Pegawai_id int
}

type Pegawai struct {
	Id        int
	Nama      string
	Username  string
	Password  string
	Email     string
	Create_at time.Time
	Delete_at time.Time
}

type Transaksi struct {
	Id              int
	Nama_produk     string
	Qty             int
	Total_transaksi int
	Create_at       time.Time
	Pelanggan_id    int
	Pegawai_id      int
}

type Pelanggan struct {
	Id         int
	Hp         string
	Nama       string
	Alamat     string
	Create_at  time.Time
	Pegawai_id int
}
type Detail_transaksi struct {
	Id              int
	Qty             int
	Total_transaksi int
	Transaksi_id    int
	Produk_id       int
}
