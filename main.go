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

	r.GET("/generate", generate)
	r.POST("/seed", changeSeed)
	r.GET("/seed", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"seed": seed})
	})

	r.POST("/contact", newContact)
	r.GET("/contact", getContact)
	r.DELETE("/contact", deleteContact)
	r.PATCH("/contact", updateContact)

	r.Run()
}
