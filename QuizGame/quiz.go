package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type quizProblem struct {
	question string
	answer   string
}

func createQuiz(data [][]string) []quizProblem {
	var myQuiz []quizProblem

	for i, line := range data {
		if i >= 0 {
			var problem quizProblem
			for j, field := range line {
				if j == 0 {
					problem.question = field
				} else if j == 1 {
					problem.answer = field
				}
			}
			myQuiz = append(myQuiz, problem)
		}
	}
	return myQuiz
}

func readFileAndConvert(file string) []quizProblem {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	myNewQuiz := createQuiz(data)

	return myNewQuiz
}

func calculateScore(userQuiz []quizProblem, answers []string) int {
	var score int

	//fmt.Println("Answers splice: ", answers[0])
	//j, err := strconv.Atoi(answers[0])
	//fmt.Println(err)
	//fmt.Println("Answers splice after converting: ", j)

	for i := range answers {

		//j, _ := strconv.Atoi(answers[i])
		//k, _ := strconv.Atoi(userQuiz[i].answer)

		//fmt.Println("Answers[i]: ", j)
		//fmt.Println("userQuiz[i].answer: ", k)

		if answers[i] == userQuiz[i].answer {
			score = score + 1
		}
	}
	return score
}

func display(userQuiz []quizProblem) {
	var answers []string

	userQuiz = readFileAndConvert("problems.csv")

	for i := range userQuiz {
		fmt.Print("Question ", i+1, ". ", userQuiz[i].question, " : ")
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		answers = append(answers, userInput)
	}

	score := calculateScore(userQuiz, answers)

	fmt.Print("Your score is: ", score, "/", len(answers))

}

func main() {
	var userQuiz []quizProblem

	display(userQuiz)
}
