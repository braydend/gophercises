package main

func main() {
	flags := setUpFlags()
	quiz := createQuizFromCsv(flags.filename)
	answers := quiz.startQuiz()
	printScore(answers)
}
