package router

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kuma0328/circle_board/infrastructure/database"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	r := &Router{
		Engine: gin.Default(),
	}
	r.setMiddleware()
	r.initHealthRouter()
	r.initCircleRouter(&database.Conn{})
	return r
}

func (r *Router) Serve() {
	// portを取得
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT is not set")
	}
	// サーバーを起動
	r.Engine.Run(fmt.Sprintf(":%s", port))
}
