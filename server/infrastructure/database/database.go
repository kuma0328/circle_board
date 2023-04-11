package database

import (
	"database/sql"
	"log"

	"github.com/kuma0328/circle_board/config"
	_ "github.com/lib/pq"
)

type Conn struct {
	DB *sql.DB
}

func initDBtable(conn *Conn) {
	query := `
		CREATE TABLE IF NOT EXISTS circle (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			image_url VARCHAR(255) NOT NULL,
			genre VARCHAR(255) NOT NULL,
			location VARCHAR(255) NOT NULL,
			member INT NOT NULL,
			twitter VARCHAR(255),
			instagram VARCHAR(255)
		)
	`
	_, err := conn.DB.Exec(query)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}
}

// データベースに接続して返す関数です
func NewDB() (*Conn, error) {
	dsn, err := config.DSN()
	if err != nil {
		log.Fatal(err)
	}

	// データベースに接続
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// データベースへの接続を確認
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// テーブルを初期化
	initDBtable(&Conn{DB: db})
	return &Conn{DB: db}, nil
}
