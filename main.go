package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// For anyone on VSCode, install Go extension

// go run .

func main() {
	//testBasics()
	//testBasicStructures()
	startServer()
}

func startServer() {
	r := gin.Default()

	// Number server
	r.GET("/generate", generate)
	r.POST("/seed", changeSeed)
	r.GET("/seed", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"seed": seed})
	})

	r.Run()
}
