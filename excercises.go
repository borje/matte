package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

func numberValidator(input string) error {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New("Invalid number")
	}
	return nil
}

func question(first int, second int, operator string) int64 {
	label := fmt.Sprintf("%d %s %d ", first, operator, second)
	prompt := promptui.Prompt{
		Label:    label,
		Validate: numberValidator,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 0
	}
	answer, _ := strconv.ParseInt(result, 10, 32)
	// fmt.Printf("You choose %q\n", result)
	return answer
}

func additionUppTill10() bool {
	for i := 0; i < 3; i++ {
		question(10, i, "-")
	}
	question(10, 2, "-")
	return true
}
