package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
)

func numberValidator(input string) error {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New("Inget giltigt tal")
	}
	return nil
}

type operator struct {
	Str string
	Op  func(int, int) int
}

type task struct {
	first  int
	second int
	op     operator
}

var plus = operator{"+", func(a int, b int) int { return a + b }}
var minus = operator{"-", func(a int, b int) int { return a - b }}

func ask(first int, second int, op operator) (bool, error) {
	label := fmt.Sprintf("%d %s %d", first, op.Str, second)
	prompt := promptui.Prompt{
		Label:    label,
		Validate: numberValidator,
	}
	result, err := prompt.Run()
	if err != nil {
		return false, err
	}
	answer, _ := strconv.ParseInt(result, 10, 0)
	return int(answer) == op.Op(first, second), nil
}

func doTasks(tasks []task) []task {
	countCorrect := 0
	var incorrect []task
	for i, task := range tasks {
		fmt.Printf("\n(%d/%d)\n", i+1, len(tasks))
		correctAnswer, err := ask(task.first, task.second, task.op)
		if err == promptui.ErrInterrupt {
			fmt.Println("Avbryter")
			break
		}
		if correctAnswer {
			fmt.Println("Rätt!")
			countCorrect++
		} else {
			fmt.Println("Fel! Rätt svar är ", task.op.Op(task.first, task.second))
			incorrect = append(incorrect, task)
			time.Sleep(time.Second)
		}
	}
	return incorrect
}

func additionUppTill(high int, count int) bool {
	var tasks []task
	for i := 0; i < count; i++ {
		a := rand.Intn(high)
		b := rand.Intn(high)
		tasks = append(tasks, task{a, b, plus})
	}
	incorrect := doTasks(tasks)
	countCorrect := len(tasks) - len(incorrect)
	fmt.Printf("Du hade %d rätt av %d\n", countCorrect, count)
	for len(incorrect) > 0 {
		incorrect = doTasks(incorrect)
		countCorrect := len(tasks) - len(incorrect)
		fmt.Printf("Du hade %d rätt av %d\n", countCorrect, count)
	}
	return true
}

func subtraktionUppTill(high int, count int) bool {
	var tasks []task
	for i := 0; i < count; i++ {
		a := rand.Intn(high)
		b := rand.Intn(high)
		if a < b {
			tasks = append(tasks, task{b, a, minus})
		} else {
			tasks = append(tasks, task{a, b, minus})
		}

	}
	incorrect := doTasks(tasks)
	countCorrect := len(tasks) - len(incorrect)
	fmt.Printf("Du hade %d rätt av %d\n", countCorrect, count)
	return true
}

func tioKompisar() bool {
	start := time.Now()
	count := 10
	var tasks []task
	for i := 0; i < count; i++ {
		b := rand.Intn(10)
		tasks = append(tasks, task{10, b, minus})
	}
	incorrect := doTasks(tasks)
	end := time.Now()
	countCorrect := len(tasks) - len(incorrect)
	fmt.Printf("Du hade %d rätt av %d\n", countCorrect, count)
	secondsPerTask := (end.Sub(start).Seconds()) / float64(len(tasks))
	fmt.Printf("Sekunder per fråga: %.1f\n", secondsPerTask)
	fmt.Println("Snabbhetspoäng: ", int(1000/float64(len(incorrect)+1)/secondsPerTask))
	time.Sleep(time.Second)
	for len(incorrect) > 0 {
		incorrect = doTasks(incorrect)
		countCorrect := len(tasks) - len(incorrect)
		fmt.Printf("Du hade %d rätt av %d\n", countCorrect, count)
	}
	return true
}
