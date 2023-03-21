package data

import (
	"database/sql"
)

type Model struct {
	conn *sql.DB
}

func (m *Model) SetSQLConnection(db *sql.DB) {
	m.conn = db
}
