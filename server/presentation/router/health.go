package router

import "github.com/gin-gonic/gin"

func initHealthRouter(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
}