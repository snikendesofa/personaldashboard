package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Load HTML files from the Web directory
	router.LoadHTMLGlob("web/html/*.html")

	// Serve static files
	//router.Static("/js", "./Web/JS")
	router.Static("/css", "./web/css")

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Start server
	err := router.Run(":8000")
	if err != nil {
		fmt.Println("Error starting server: " + err.Error())
	}
}
