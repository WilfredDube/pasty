package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from Pasty",
	})
}

func snippetView(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "paste not specified....",
		})
		return
	}

	log.Printf("id = %+v", id)
	c.JSON(200, gin.H{
		"message": "Display a specific paste...",
	})
}

func snippetCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create a new paste...",
	})
}
