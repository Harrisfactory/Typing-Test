package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	//Setup regular routing
	site := router.Group("/")
	{
		site.GET("/start", StartTypingTest)
	}

	//endpoints here

	router.Run(":4000") // 8080/ping
}

func StartTypingTest(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "typing test not yet implemented",
	})
}
