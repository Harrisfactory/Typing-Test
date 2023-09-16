package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"typing_test/src/typing"
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

	//typing test testing
	tt := typing.TypingTest{
		Time: 60,
	}

	tt.GetWords()
	tt.PrintWords()

	router.Run(":4000") // 8080/ping
}

func StartTypingTest(c *gin.Context) {

	tt := typing.TypingTest{
		Time: 60,
	}

	tt.GetWords()
	tt.PrintWords()

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": tt.Words,
	})
}
