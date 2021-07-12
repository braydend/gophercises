package main

type Question struct {
	Prompt string
	Answer string
}

type Quiz struct {
	Questions []Question
}

func (q *Quiz) addQuestion(prompt string, answer string) *Quiz {
	q.Questions = append(q.Questions, Question{prompt, answer})

	return q
}

func parseQuiz(filename string) *Quiz{
	reader := getCsvReader(filename)
	data := readCsv(reader)
	quiz := &Quiz{}

	for _, row := range data {
		quiz.addQuestion(row[0], row[1])
	}

	return quiz
}