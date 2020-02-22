package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Sorting(lines []string, opt *Options) ([]string, error) {
	var err error

	if opt.nFlag {
		lines, err = IntSorting(lines, opt)
	} else {
		lines, err = StringSorting(lines, opt)
	}

	if err != nil {
		return nil, err
	}

	if opt.uFlag {
		lines = DeleteDuplicates(lines, opt)
	}

	if opt.oFlag != "" {
		return lines, writeIntoFile(lines, *o)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return lines, nil
}

func StringSorting(lines []string, opt *Options) (result []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	if opt.fFlag {
		sort.Slice(lines, func(i, j int) bool {
			if opt.kFlag < 0 {
				return strLess(strings.ToUpper(lines[i]), strings.ToUpper(lines[j]), opt.rFlag)
			}

			lhs := strings.Split(lines[i], " ")[opt.kFlag]
			rhs := strings.Split(lines[j], " ")[opt.kFlag]
			return strLess(strings.ToUpper(lhs), strings.ToUpper(rhs), opt.rFlag)
		})
		return lines, nil
	}

	sort.Slice(lines, func(i, j int) bool {
		if opt.kFlag < 0 {
			return strLess(lines[i], lines[j], opt.rFlag)
		}

		lhs := strings.Split(lines[i], " ")[opt.kFlag]
		rhs := strings.Split(lines[j], " ")[opt.kFlag]
		return strLess(lhs, rhs, opt.rFlag)
	})
	return lines, nil
}

func IntSorting(lines []string, opt *Options) (result []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	sort.Slice(lines, func(i, j int) bool {
		var lhs, rhs int

		if opt.kFlag < 0 {
			lhs, err = strconv.Atoi(lines[i])
			rhs, err = strconv.Atoi(lines[j])
		} else {
			lhs, err = strconv.Atoi(strings.Split(lines[i], " ")[opt.kFlag])
			rhs, err = strconv.Atoi(strings.Split(lines[j], " ")[opt.kFlag])
		}

		return intLess(lhs, rhs, opt.rFlag)
	})

	return lines, err
}

func strLess(lhs, rhs string, reverse bool) bool {
	if reverse {
		return lhs > rhs
	}

	return lhs < rhs
}

func intLess(lhs, rhs int, reverse bool) bool {
	if reverse {
		return lhs > rhs
	}

	return lhs < rhs
}

func DeleteDuplicates(lines []string, opt *Options) []string {
	var result []string
	lineExist := make(map[string]bool)

	for _, line := range lines {
		key := line
		if opt.fFlag {
			key = strings.ToUpper(key)
		}
		if _, flag := lineExist[key]; !flag {
			lineExist[key] = true
			result = append(result, line)
		}
	}

	return result
}
