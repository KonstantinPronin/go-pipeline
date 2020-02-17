package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Sorting(lines []string) ([]string, error) {
	var err error

	if *n {
		lines, err = IntSorting(lines)
	} else {
		lines, err = StringSorting(lines, *k)
	}

	if err != nil {
		return nil, err
	}

	if *u {
		lines = DeleteDuplicates(lines)
	}

	if *r {
		lines = Reverse(lines)
	}

	if *o != "" {
		return lines, writeIntoFile(lines, *o)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return lines, nil
}

func StringSorting(lines []string, column int) (result []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	if *f {
		sort.Slice(sort.StringSlice(lines), func(i, j int) bool {
			if column < 0 {
				return strings.ToUpper(lines[i]) < strings.ToUpper(lines[j])
			}

			lhs := strings.Split(lines[i], " ")[column]
			rhs := strings.Split(lines[j], " ")[column]
			return strings.ToUpper(lhs) < strings.ToUpper(rhs)
		})
		return lines, nil
	}

	sort.Slice(sort.StringSlice(lines), func(i, j int) bool {
		if column < 0 {
			return lines[i] < lines[j]
		}

		lhs := strings.Split(lines[i], " ")[column]
		rhs := strings.Split(lines[j], " ")[column]
		return lhs < rhs
	})
	return lines, nil
}

func IntSorting(lines []string) ([]string, error) {
	var numbers []int
	var result []string

	for _, str := range lines {
		number, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	sort.Ints(numbers)
	for _, number := range numbers {
		str := strconv.Itoa(number)
		result = append(result, str)
	}

	return result, nil
}

func Reverse(lines []string) []string {
	for i := 0; i < len(lines)/2; i++ {
		j := len(lines) - 1 - i
		lines[i], lines[j] = lines[j], lines[i]
	}

	return lines
}

func DeleteDuplicates(lines []string) []string {
	var result []string
	lineExist := make(map[string]bool)

	for _, line := range lines {
		key := line
		if *f {
			key = strings.ToUpper(key)
		}
		if _, flag := lineExist[key]; !flag {
			lineExist[key] = true
			result = append(result, line)
		}
	}

	return result
}
