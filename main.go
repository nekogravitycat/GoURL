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

	router.GET("/:shortened", urlRedirect)

	router.GET("/admin", createEntryView)
	router.POST("/admin", createEntrySubmit)

	err := router.Run(":8080")
	if err != nil {
		panic("[ERROR] failed to start Gin server, error: " + err.Error())
	}
}
