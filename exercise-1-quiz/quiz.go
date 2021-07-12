package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Question struct {
	Prompt string
	Answer string
}

type Quiz struct {
	Questions []Question
}

type Answer struct {
	question Question
	isCorrect bool
}

func (quiz *Quiz) addQuestion(prompt string, answer string) *Quiz {
	quiz.Questions = append(quiz.Questions, Question{prompt, answer})

	return quiz
}

func (question *Question) answerQuestion() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(question.Prompt)

	input, _, err := reader.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	return string(input) == question.Answer
}

func (quiz *Quiz) startQuiz() (answers []*Answer) {
	for _, question := range quiz.Questions {
		answers = append(answers, &Answer{question, question.answerQuestion()})
	}

	return answers
}

func printScore(answers []*Answer) {
	score := 0
	for _, answer := range answers {
		if answer.isCorrect {
			score++
		}
	}

	fmt.Printf("You scored %d/%d!\n", score, len(answers))
}

func createQuizFromCsv(filename string) *Quiz{
	data := readCsv(filename)
	quiz := &Quiz{}

	for _, row := range data {
		quiz.addQuestion(row[0], row[1])
	}

	return quiz
}