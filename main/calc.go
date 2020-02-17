package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strconv"
)

var priority = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func ConvertToRpn(expression []string) []string {
	var result []string
	await := stack.New()

	for _, symbol := range expression {
		if _, err := strconv.Atoi(symbol); err == nil {
			result = append(result, symbol)
			continue
		}

		top := await.Peek()

		switch {
		case top == nil || symbol == "(":
			await.Push(symbol)
		case symbol == ")":
			for operator := await.Pop(); operator != nil && operator.(string) != "("; {
				result = append(result, operator.(string))
				operator = await.Pop()
			}
		case priority[symbol] > priority[top.(string)]:
			await.Push(symbol)
		case priority[symbol] <= priority[top.(string)]:
			result = append(result, top.(string))
			await.Pop()
			await.Push(symbol)
		}
	}

	for operator := await.Pop(); operator != nil; {
		result = append(result, operator.(string))
		operator = await.Pop()
	}

	return result
}

func Calculate(expression []string) (int, error) {
	result := stack.New()

	for _, symbol := range expression {
		if number, err := strconv.Atoi(symbol); err == nil {
			result.Push(number)
			continue
		}

		rhs := result.Pop()
		lhs := result.Pop()

		if lhs == nil || rhs == nil {
			return 0, fmt.Errorf("wrong input expression")
		}

		switch symbol {
		case "+":
			result.Push(lhs.(int) + rhs.(int))
		case "-":
			result.Push(lhs.(int) - rhs.(int))
		case "*":
			result.Push(lhs.(int) * rhs.(int))
		case "/":
			result.Push(lhs.(int) / rhs.(int))
		default:
			return 0, fmt.Errorf("wrong input expression")
		}
	}

	return result.Peek().(int), nil
}
