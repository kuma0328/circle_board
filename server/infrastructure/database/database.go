package database

import (
	"database/sql"

	"github.com/kuma0328/circle_board/config"
)

// データベースに接続して返す関数です
func NewDB() (*sql.DB, error) {
	dsn, driver := config.DSN()

	// データベースに接続
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	// データベースへの接続を確認
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}