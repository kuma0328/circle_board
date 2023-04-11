package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kuma0328/circle_board/infrastructure/database"
)

func InitRouter(conn *database.Conn) *gin.Engine {
	r := gin.Default()
	

	// 各ルーティングを初期化	
	initHealthRouter(r)
	initCircleRouter(r, conn)
	return r
}