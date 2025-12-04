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
		ctx.HTML(http.StatusOK, "index.html", gin.H {
			"content": "hello world",
		})
	})

	router.Run("localhost:6969")
}
