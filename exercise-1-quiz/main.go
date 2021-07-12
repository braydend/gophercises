package main

import (
	"fmt"
)

func main() {
	flags := setUpFlags()
	quiz := parseQuiz(flags.filename)

	fmt.Print(quiz)
}
