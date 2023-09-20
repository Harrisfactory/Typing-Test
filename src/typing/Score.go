package typing

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Score struct {
	Name  string
	Score *float64
}

func (s *Score) TopScores() []Score {
	var scores []Score

	//connect do db
	db, err := sql.Open("sqlite3", "typing-test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//fetch top 10 scores
	rows, err := db.Query("SELECT name, score FROM scores ORDER BY score DESC LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//add them to scores struct
	for rows.Next() {
		var score Score
		if err := rows.Scan(&score.Name, &score.Score); err != nil {
			log.Fatal(err)
		}
		scores = append(scores, score)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return scores
}
