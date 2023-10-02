package main

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// for get method
func createEntryView(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{
		"title": "Create an entry",
	})
}

// for post method
func createEntrySubmit(c *gin.Context) {
	shortened := strings.TrimSpace(c.PostForm("shortened"))
	destination := strings.TrimSpace(c.PostForm("destination"))
	override := strings.TrimSpace(c.PostForm("override"))

	if shortened == "" {
		shortened = randomString(6)
	}

	override_bool := override != ""
	status, message := createURL(shortened, destination, override_bool)

	switch status {
	case "successful":
		short_url := "t.gravitycat.tw/" + shortened
		c.HTML(http.StatusOK, "create.html", gin.H{
			"status": "successful",
			"detail": short_url,
		})
	case "override-confirm":
		c.HTML(http.StatusAccepted, "create.html", gin.H{
			"status":      "interrupted",
			"override":    "1",
			"destination": destination,
			"shortened":   shortened,
			"detail":      message,
		})
	case "failed":
		c.HTML(http.StatusBadRequest, "create.html", gin.H{
			"status": "failed",
			"detail": message,
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unknown operation status",
		})
	}
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(length int) string {
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
