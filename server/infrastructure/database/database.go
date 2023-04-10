package database

import (
	"database/sql"

	"github.com/kuma0328/circle_board/config"
)

func NewDB() (*sql.DB, error) {
	dsn, driver := cofig.DSN()
}