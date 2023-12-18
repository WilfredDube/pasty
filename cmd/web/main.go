package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()
	router := gin.Default()

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	router.LoadHTMLFiles(files...)
	router.Static("/assets", "./ui/static/")

	router.GET("/", home)
	router.GET("/paste/view", snippetView)
	router.POST("/paste/create", snippetCreate)

	router.Run(":4000")
}
