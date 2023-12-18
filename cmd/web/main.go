package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./ui/html/pages/*")
	router.Static("/assets", "./ui/static/")

	router.GET("/", home)
	router.GET("/paste/view", snippetView)
	router.POST("/paste/create", snippetCreate)

	router.Run(":4000")
}
