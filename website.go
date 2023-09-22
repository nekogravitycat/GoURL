package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	loadURLTable()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/admin")
	})
	router.GET("/:url", urlRedirect)

	router.GET("/admin", createEntryPage)
	router.POST("/admin", createEntry)

	router.Run(":8080")
}
