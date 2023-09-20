package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"typing_test/src/typing"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	//routes
	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "Main website",
	//	})
	//})

	//start test
	router.GET("/", StartTypingTest)

	//submit to score table
	router.POST("/submit-score", SubmitScore)

	router.Run(":4000") // 8080/ping
}

func SubmitScore(c *gin.Context) {
	//grab displayName from formdata
	displayName := c.PostForm("display_name")
	score := c.PostForm("final-score")

	//insert into db
	//db
	db, err := sql.Open("sqlite3", "typing-test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a 'scores' table if not exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS scores (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            score INT
        )
    `)
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO scores (name, score) VALUES (?,?)", displayName, score)
	if err != nil {
		panic(err)
	}

	c.Redirect(http.StatusFound, "/")
}

func StartTypingTest(c *gin.Context) {

	tt := typing.TypingTest{
		Time: 60,
	}

	s := typing.Score{}

	tt.HarvestWords()
	//tt.PrintWords()

	c.HTML(http.StatusOK, "typing_page.html", gin.H{
		"title":  "le typing test",
		"words":  tt.Words,
		"Scores": s.TopScores(),
	})

}
