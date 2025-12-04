package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Defl8/what-is-wyatt-doing/internal/requests"
	"github.com/Defl8/what-is-wyatt-doing/internal/github"
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

	reqHandler := requests.NewRequestHandler(nil, os.Getenv("GITHUB_TOKEN"))

	router.GET("/stinky", func(ctx *gin.Context) {
		events, err := reqHandler.GetPublicUserEvents("Defl8")
		if err != nil {
			log.Fatal("User event data request failed.")
		}

		displayEvents := []github.DisplayEvent{}
		for _, e := range *events {
			displayEvents = append(displayEvents, e.Display())
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"content": displayEvents,
		})
	})

	router.Run("localhost:6969")
}
