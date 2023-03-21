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

func (m *Model) TambahPegawai(newPegawai Pegawai) error {
	res, err := m.conn.Exec("INSERT INTO pegawai (nama, username, password,email) values(?,?,?,?)", newPegawai.Nama, newPegawai.Username, newPegawai.Password, newPegawai.Email)

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

	m.conn.Close()

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

	m.conn.Close()

	return nil
}

func (m Model) GetAllPegawai() ([]Pegawai, error) {
	listPegawai := []Pegawai{}
	rows, err := m.conn.Query("SELECT idpegawai, nama, username FROM pegawai")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var newPegawai = Pegawai{}
		if err := rows.Scan(&newPegawai.Id, &newPegawai.Nama, &newPegawai.Username); err != nil {
			fmt.Println(err)
			return nil, err
		}
		listPegawai = append(listPegawai, newPegawai)
	}
	m.conn.Close()

	return listPegawai, nil
}
