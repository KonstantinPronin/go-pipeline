package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var validChar = map[rune]struct{}{
	'+': {},
	'-': {},
	'*': {},
	'/': {},
	'(': {},
	')': {},
}

func Parse(expression string) ([]string, error) {
	var result []string
	var err error
	var number string

	for _, runeValue := range expression {
		if _, err := strconv.Atoi(string(runeValue)); err == nil {
			number += string(runeValue)
		} else if _, ok := validChar[runeValue]; ok {
			if number != "" {
				result = append(result, number)
			}
			result = append(result, string(runeValue))
			number = ""
		} else {
			err = fmt.Errorf("invalid input")
		}
	}

	if number != "" {
		result = append(result, number)
	}

	return result, err
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}

	expression, err := Parse(os.Args[1])
	if err != nil {
		log.Fatal("Invalid input")
	}

	expression = ConvertToRpn(expression)
	result, err := Calculate(expression)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
