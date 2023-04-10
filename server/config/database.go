package config

import (
	"fmt"
	"os"
)

// データベースに接続するための関数
func DSN() (string, string) {
	// 環境変数からデータベースの情報を取得
	SQLDriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// DSNを作成
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return dsn, SQLDriver
}