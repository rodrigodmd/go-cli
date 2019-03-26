package prompt

import (
	"fmt"
	"os"

	"gopkg.in/AlecAivazis/survey.v1"
)

func Input(label string) string {
	response := ""
	prompt := &survey.Input{
		Message: label,
	}
	err := survey.AskOne(prompt, &response, nil)

	if err != nil {
		fmt.Println("Canceled")
		os.Exit(1)
		return ""
	}

	return response
}

func Select(label string, items []string) string {
	response := ""
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}
	err := survey.AskOne(prompt, &response, nil)

	if err != nil {
		fmt.Println("Canceled")
		os.Exit(1)
		return ""
	}

	return response
}

func MultiSelect(label string, items []string) []string {
	response := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: items,
	}
	err := survey.AskOne(prompt, &response, nil)

	if err != nil {
		fmt.Println("Canceled")
		os.Exit(1)
		return []string{}
	}

	return response
}

func Confirm(label string) bool {
	response := false
	prompt := &survey.Confirm{
		Message: label,
	}
	err := survey.AskOne(prompt, &response, nil)

	if err != nil {
		fmt.Println("Canceled")
		os.Exit(1)
		return false
	}

	return response
}
