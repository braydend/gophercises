package main

func main() {
	flags := setUpFlags()
	createQuizFromCsv(flags.filename).startQuiz(flags.timeLimit)
}
