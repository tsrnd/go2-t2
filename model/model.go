package model

import (
	"database/sql"
	"time"
)

// Model general struct
type Model struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

//DBCon dbcon
var DBCon *sql.DB

//SetDatabase return DBCon
func SetDatabase(database *sql.DB) {
	DBCon = database
}
