package main

import (
	"fmt"

	"example.com/quiz/reader"
)

func main() {
	fmt.Println("Welcome to math quizziez.")
	var userInput string
	reader.CompareNumbers(&userInput)
}
