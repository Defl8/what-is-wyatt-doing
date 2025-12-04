package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Load Templates
	router.LoadHTMLGlob("web/template/**")

	// Load Static Content
	router.Static("/static", "web/static/")


	router.GET("/stinky", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H {
			"msg": "Hello, World",
		})
	})

	router.Run("localhost:6969")
}
