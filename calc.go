package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("")
	fmt.Print(">")

	for scanner.Scan() {
		text := scanner.Text()
		var result int64

		// Format Input
		expressionString := formatExpression(text)

		// Parse Argument
		num1, num2, operator, err := parseArgs(expressionString)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Caculation
		result, err = calc2Number(num1, num2, operator)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(expressionString+" =", result)
		fmt.Print(">")
	}
}

func formatExpression(text string) string {
	// Trim Space
	formatedText := strings.TrimSpace(text)
	// Remove multiple space
	regex := regexp.MustCompile(" +")
	formatedText = regex.ReplaceAllString(formatedText, "")
	// Add Space before Operator [Ex: 9*5 => 9 *5]
	regex = regexp.MustCompile("(\\d)([*+-/])")
	formatedText = regex.ReplaceAllString(formatedText, "$1 $2")
	// Add Space after Operator [Ex: 9*5 => 9* 5]
	regex = regexp.MustCompile("([*+-/])(\\d)")
	formatedText = regex.ReplaceAllString(formatedText, "$1 $2")

	return formatedText
}

func parseArgs(text string) (num1, num2 int64, operator string, err error) {
	element := strings.Split(text, " ")
	if len(element) != 3 {
		return 0, 0, "", errors.New("Invalid arguments!")
	}

	operator = element[1]
	num1, err = strconv.ParseInt(element[0], 10, 64)
	if err != nil {
		return 0, 0, "", errors.New("Invalid argument 1!")
	}
	num2, err = strconv.ParseInt(element[2], 10, 64)
	if err != nil {
		return 0, 0, "", errors.New("Invalid argument 2!")
	}
	return
}

func calc2Number(num1, num2 int64, operator string) (result int64, err error) {

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return 0.0, errors.New("Err : Divided by zero!")
		}
		result = num1 / num2
	default:
		return 0.0, errors.New("Err : Invalid Operator!")
	}
	return
}
