package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"typing_test/src/typing"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	//routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	//start test
	router.GET("/start", StartTypingTest)

	router.Run(":4000") // 8080/ping
}

func StartTypingTest(c *gin.Context) {

	tt := typing.TypingTest{
		Time: 60,
	}

	tt.HarvestWords()
	tt.PrintWords()

	c.HTML(http.StatusOK, "typing_page.html", gin.H{
		"title": "le typing test",
		"words": tt.Words,
	})

}
