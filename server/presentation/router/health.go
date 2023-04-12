package router

import "github.com/gin-gonic/gin"

func (r *Router) initHealthRouter() {
	r.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
}