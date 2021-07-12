package main

func main() {
	flags := setUpFlags()
	quiz := parseQuiz(flags.filename)
	answers := quiz.startQuiz()
	printScore(answers)
}
