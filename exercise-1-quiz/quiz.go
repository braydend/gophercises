package main

import (
	"fmt"
)

type Question struct {
	Prompt string
	Answer string
}

type Quiz struct {
	Questions []Question
}

func (quiz *Quiz) addQuestion(prompt string, answer string) *Quiz {
	quiz.Questions = append(quiz.Questions, Question{prompt, answer})

	return quiz
}

func readAnswer(answerChannel chan string) {
	answerChannel <- string(readInput())
}

func verifyAnswer(input string, answer string) bool {
	return input == answer
}

func (quiz *Quiz) startQuiz(timeLimit int) {
	answerChannel := make(chan string)
	timer := startTimer(timeLimit)
	correctAnswers := 0

	quizLoop:
	for _, question := range quiz.Questions {
		fmt.Println(question.Prompt)
		go readAnswer(answerChannel)

		select{
			case answer := <- answerChannel:
				if verifyAnswer(answer, question.Answer) {
					correctAnswers++
				}
				break

			case <- timer.C:
				break quizLoop
		}
	}


	printScore(correctAnswers, len(quiz.Questions))
}

func printScore(correctAnswers int, totalQuestions int) {
	fmt.Printf("You scored %d out of %d!\n", correctAnswers, totalQuestions)
}

func createQuizFromCsv(filename string) *Quiz{
	data := readCsv(filename)
	quiz := &Quiz{}

	for _, row := range data {
		quiz.addQuestion(row[0], row[1])
	}

	return quiz
}