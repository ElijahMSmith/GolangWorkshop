package main

// Show that you can also run this through Cobra as a command line utility
import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	seed int = 150
)

func generate(c *gin.Context) {
	countStr := c.Query("count")
	if countStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Count query parameter not specified",
		})
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil || count < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid count query parameter",
		})
		return
	}

	values := make([]int, count)
	for i := 0; i < count; i++ {
		values[i] = rand.Int()
	}

	c.JSON(http.StatusOK, gin.H{
		"values": values,
	})
}

func changeSeed(c *gin.Context) {
	seedStr := c.Param("seed")
	if seedStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Seed parameter not specified",
		})
		return
	}

	newSeed, err := strconv.Atoi(seedStr)
	if err != nil || seed < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid seed parameter (must be integer > 0)",
		})
		return
	}

	seed = newSeed
	rand.Seed(int64(seed))
	c.JSON(http.StatusAccepted, gin.H{})
}
