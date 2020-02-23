package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var validChar = map[rune]struct{}{
	'+': {},
	'-': {},
	'*': {},
	'/': {},
	'(': {},
	')': {},
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}

	argument := strings.Join(os.Args[1:], "")
	expression, err := Parse(argument)
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
