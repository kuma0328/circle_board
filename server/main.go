package main

import (
	"fmt"
	"log"

	"github.com/kuma0328/circle_board/config"
	"github.com/kuma0328/circle_board/infrastructure/database"
	"github.com/kuma0328/circle_board/presentation/router"
)


func main() {
	fmt.Println(config.DSN())
	// NewDB()を呼び出してデータベースに接続
	conn, err := database.NewDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer conn.DB.Close()

	// routerを初期化
	r := router.NewRouter()
	
	// Routerの起動
	r.Serve()
}
