package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kuma0328/circle_board/infrastructure/database"
	"github.com/kuma0328/circle_board/infrastructure/persistance"
	"github.com/kuma0328/circle_board/presentation/handler"
	"github.com/kuma0328/circle_board/usecase"
)

func initCircleRouter(r *gin.Engine, conn *database.Conn) {
	repo := persistance.NewCircleRepository(conn)
	uc := usecase.NewCircleUsecase(repo)
	h := handler.NewCircleHandler(uc)

	g := r.Group("/circle")
	g.GET("/all", h.GetAllCircle)
	g.GET("/:id", h.GetCircleById)
	g.POST("/create", h.CreateCircle)
}