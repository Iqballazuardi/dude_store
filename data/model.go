package data

import (
	"database/sql"
	"errors"
	"fmt"
)

type Model struct {
	conn *sql.DB
}

func (m *Model) SetSQLConnection(db *sql.DB) {
	m.conn = db
}
func (m *Model) Login(username string, password string) (*Pegawai, error) {
	var pegawai = Pegawai{}

	err := m.conn.QueryRow("SELECT nama, idpegawai FROM pegawai where username = ? AND password = ?",
		username, password).Scan(&pegawai.Nama, &pegawai.Id)

	if err != nil {
		return nil, err
	}
	// m.conn.Close()
	return &pegawai, nil
}

func (m *Model) TambahPegawai(Pegawai) error {
	var pegawai = Pegawai{}
	fmt.Print("Masukkan Nama Pegawai: ")
	fmt.Scanln(&pegawai.Nama)
	fmt.Print("Masukkan Username Pegawai: ")
	fmt.Scanln(&pegawai.Username)
	fmt.Print("Masukkan Password Pegawai: ")
	fmt.Scanln(&pegawai.Password)
	fmt.Print("Masukkan Email Pegawai: ")
	fmt.Scanln(&pegawai.Email)
	res, err := m.conn.Exec("INSERT INTO pegawai (nama, username, password,email) values(?,?,?,?)",
		pegawai.Nama, pegawai.Username, pegawai.Password, pegawai.Email)

	if err != nil {
		fmt.Println(err)
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if aff <= 0 {
		return errors.New("terjadi sebuah masalah pada sistem")
	}

	// m.conn.Close()

	return nil
}
func (m *Model) DeletePegawai(Email string) error {
	res, err := m.conn.Exec("DELETE FROM pegawai WHERE email = ?", Email)
	if err != nil {
		fmt.Println(err)
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if aff <= 0 {
		return errors.New("terjadi sebuah masalah pada sistem")
	}

	// m.conn.Close()

	return nil
}

func (m Model) GetAllPegawai() ([]Pegawai, error) {
	listPegawai := []Pegawai{}
	rows, err := m.conn.Query("SELECT idpegawai, nama, username, email, create_at FROM pegawai")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// createdAt, err := time.Parse("2006-01-02 15:04:05", string(newPegawai.Create_at))

	for rows.Next() {
		var newPegawai = Pegawai{}
		if err := rows.Scan(&newPegawai.Id, &newPegawai.Nama, &newPegawai.Username, &newPegawai.Email, &newPegawai.Create_at); err != nil {
			fmt.Println(err)
			return nil, err
		}

		listPegawai = append(listPegawai, newPegawai)
	}
	// m.conn.Close()

	return listPegawai, nil
}
func (m *Model) TambahProduk(newProduk Produk) error {
	res, err := m.conn.Exec("INSERT INTO produk (nama, keterangan, stok , harga, pegawai_idpegawai) values(?,?,?,?,?)", newProduk.Nama, newProduk.Keterangan, newProduk.Stok, newProduk.Harga, newProduk.Pegawai_id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if aff <= 0 {
		return errors.New("terjadi sebuah masalah pada sistem")
	}

	// m.conn.Close()

	return nil
}
func (m *Model) TambahPelanggan(newPelanggan Pelanggan) error {
	res, err := m.conn.Exec("INSERT INTO pelanggan (nama, hp, alamat, pegawai_idpegawai) values(?,?,?,?)", newPelanggan.Nama, newPelanggan.Hp, newPelanggan.Alamat, newPelanggan.Pegawai_id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if aff <= 0 {
		return errors.New("terjadi sebuah masalah pada sistem")
	}

	// m.conn.Close()

	return nil
}

func (m *Model) GetProdukById(id int) (*Produk, error) {
	row := m.conn.QueryRow("SELECT idproduk, nama, keterangan, stok, harga  FROM produk WHERE idproduk = ?", id)

	var produk Produk
	err := row.Scan(&produk.Id, &produk.Nama, &produk.Keterangan, &produk.Stok, &produk.Harga)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("produk tidak ditemukan")
		} else {
			return nil, err
		}
	}

	return &produk, nil
}

func (m *Model) UpdateProduk(id int, updatedProduk Produk) error {
	tx, err := m.conn.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = m.GetProdukById(id)
	if err != nil {
		return err
	}
	res, err := tx.Exec("UPDATE produk SET nama = ?, keterangan = ?, stok = ?, harga = ? WHERE idproduk = ?",
		updatedProduk.Nama, updatedProduk.Keterangan, updatedProduk.Stok, updatedProduk.Harga, id)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if aff <= 0 {
		return errors.New("tidak ada produk yang diupdate")
	}

	return nil
}
