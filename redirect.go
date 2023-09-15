package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func urlRedirect(c *gin.Context) {
	url := c.Param("url")

	destination, ok := urlTable[url]

	if ok {
		c.Redirect(http.StatusTemporaryRedirect, destination)
	} else {
		c.HTML(http.StatusNotFound, "notfound.html", gin.H{
			"title": "NOT FOUND",
		})
	}
}
