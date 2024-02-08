package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World")
}

func quoteHandler(ctx *gin.Context) {
	quotes := []string{
		"Believe you can and you're halfway there.",
		"Don't watch the clock; do what it does. Keep going.",
		"Start where you are. Use what you have. Do what you can.",
		"Success is the sum of small efforts, repeated day in and day out.",
		"Act as if what you do makes a difference. It does.",
		"The only way to do great work is to love what you do.",
		"Believe in yourself! Have faith in your abilities!",
		"Your talent is God's gift to you. What you do with it is your gift back to God.",
	}

	rand.Seed(time.Now().UnixNano())
	randomQuote := quotes[rand.Intn(len(quotes))]
	ctx.JSON(http.StatusOK, gin.H{"quote": randomQuote})
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/", handler)
	router.GET("/quote", quoteHandler)

	fmt.Println("Go REST service started...")
	router.Run("0.0.0.0:9009")
}
