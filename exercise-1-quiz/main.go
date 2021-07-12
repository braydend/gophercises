package main

import (
	"fmt"
)

func main() {
	quiz := parseQuiz("addition.csv")

	fmt.Print(quiz)
}
