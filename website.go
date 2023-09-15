package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	loadURLTable()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	//router.GET("/")
	router.GET("/:url", urlRedirect)

	//router.GET("/admin", createEntryPage)
	//router.POST("/admin", createEntry)

	router.Run(":8080")
}
