package typing

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type TypingTest struct {
	Time  int
	Wpm   int
	Words [250]string
}

func (tt *TypingTest) HarvestWords() {
	file, err := os.Open("word_banks/top_1000_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i < 249 && scanner.Scan(); i++ {
		line := scanner.Text()
		tt.Words[i] = line
	}
	tt.RandomizeWords()
}

func (tt *TypingTest) RandomizeWords() {
	//from last element to the first
	for i := len(tt.Words) - 1; i > 0; i-- {
		//generate random index
		j := rand.Intn(i + 1)

		tt.Words[i], tt.Words[j] = tt.Words[j], tt.Words[i]
	}
}

func (tt *TypingTest) PrintWords() {
	for i := 0; i < 249; i++ {
		fmt.Print(tt.Words[i] + " ")
	}
}
