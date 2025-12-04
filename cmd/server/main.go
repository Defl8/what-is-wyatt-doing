package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil && os.Getenv("ENV") != "PROD" {
		log.Fatal("Failed to load .env file.")
	}
}

func main() {
	router := gin.Default()

	// Load Templates
	router.LoadHTMLGlob("web/template/**")

	// Load Static Content
	router.Static("/static", "web/static/")

	router.GET("/stinky", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"content": "hello world",
		})
	})

	router.Run("localhost:6969")
}
