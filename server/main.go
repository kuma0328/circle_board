package main

import (
	"fmt"
	"log"
	"os"

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
	router := router.InitRouter(conn)

	// portを取得
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT is not set")
	}
	// サーバーを起動
	router.Run(fmt.Sprintf(":%s", port))
}
