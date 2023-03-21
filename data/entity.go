package data

import "time"

type Produk struct {
	id         int
	nama       string
	keterangan string
	stok       int
	harga      int
	create_at  time.Time
	update_at  time.Time
	pegawai_id int
}

type Pegawai struct {
	id        int
	nama      string
	username  string
	password  string
	create_at time.Time
	delete_at time.Time
}

type Transaksi struct {
	id              int
	nama_produk     string
	qty             int
	total_transaksi int
	create_at       time.Time
	pelanggan_id    int
	pegawai_id      int
}

type Pelanggan struct {
	id         int
	hp         string
	nama       string
	alamat     string
	create_at  time.Time
	pegawai_id int
}
type Detail_transaksi struct {
	id              int
	qty             int
	total_transaksi int
	transaksi_id    int
	produk_id       int
}
