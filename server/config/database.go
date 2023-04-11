package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// 環境変数を読み込む
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

// データベースに接続するための関数
func DSN() (string, error) {

	err := LoadEnv()
	if err != nil {
		return "", err
	}

	// 環境変数からデータベースの情報を取得
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return "", fmt.Errorf("DB_HOST is not set")
	}
	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return "", fmt.Errorf("DB_PORT is not set")
	}
	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return "", fmt.Errorf("DB_USER is not set")
	}
	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return "", fmt.Errorf("DB_PASSWORD is not set")
	}
	dbname, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return "", fmt.Errorf("DB_NAME is not set")
	}

	// DSNを作成
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return dsn, nil
}
