package main

import (
	"fmt"
	"testing"
)

func fString(input []string) string {
	var formatStr string
	for _, str := range input {
		formatStr += fmt.Sprintf("%s, ", str)
	}
	return formatStr
}

func TestSorting(t *testing.T) {
	opt := Options{
		fFlag: true,
		uFlag: true,
		rFlag: true,
		kFlag: 1,
	}
	input := []string{"hhh bbb", "bbb aaa", "zzz AAA", "ooo yyy", "BBB AAA"}
	expected := []string{"ooo yyy", "hhh bbb", "bbb aaa", "zzz AAA"}

	output, err := Sorting(input, &opt)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	for i, str := range expected {
		if str != output[i] {
			t.Errorf("Call: Sorting(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestSortingAllFlags(t *testing.T) {
	opt := Options{
		fFlag: true,
		uFlag: true,
		rFlag: true,
		oFlag: "",
		nFlag: true,
		kFlag: 1,
	}
	input := []string{"1 3 2", "2 1 3", "3 2 1", "1 3 2"}
	expected := []string{"1 3 2", "3 2 1", "2 1 3"}

	output, err := Sorting(input, &opt)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	for i, str := range expected {
		if str != output[i] {
			t.Errorf("Call: Sorting(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestStringSorting(t *testing.T) {
	opt := Options{kFlag: -1}
	input := []string{"hhh bbb", "bbb aaa", "zzz AAA"}
	expected := []string{"bbb aaa", "hhh bbb", "zzz AAA"}

	output, _ := StringSorting(input, &opt)

	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: StringSorting(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestStringSortingInvalidColumn(t *testing.T) {
	opt := Options{kFlag: 10}
	input := []string{"hhh bbb", "bbb aaa", "zzz AAA"}

	_, err := StringSorting(input, &opt)

	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestStringSortingByColumn(t *testing.T) {
	opt := Options{kFlag: 1}
	input := []string{"hhh bbb", "bbb aaa", "zzz AAA"}
	expected := []string{"zzz AAA", "bbb aaa", "hhh bbb"}

	output, _ := StringSorting(input, &opt)

	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: StringSorting(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestStringSortingByColumnIgnoreLetterCase(t *testing.T) {
	opt := Options{fFlag: true, kFlag: 1}
	input := []string{"hhh bbb", "zzz AAA", "bbb aaa"}
	expected := []string{"zzz AAA", "bbb aaa", "hhh bbb"}

	output, _ := StringSorting(input, &opt)

	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: StringSorting(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestStringSortingIgnoreLetterCase(t *testing.T) {
	opt := Options{fFlag: true, kFlag: -1}
	input := []string{"abc", "ABC"}
	expected := []string{"abc", "ABC"}

	output, _ := StringSorting(input, &opt)

	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: StringSorting(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestIntSortingWithError(t *testing.T) {
	opt := Options{kFlag: -1}
	input := []string{"abc", "qwe"}

	_, err := IntSorting(input, &opt)

	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestIntSortingNoError(t *testing.T) {
	opt := Options{kFlag: -1}
	input := []string{"3", "2", "1"}
	expected := []string{"1", "2", "3"}

	output, err := IntSorting(input, &opt)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: Revers(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestReverse(t *testing.T) {
	opt := Options{rFlag: true}
	input := []string{"1", "2", "3"}
	expected := []string{"3", "2", "1"}

	output, err := Sorting(input, &opt)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: Revers(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestDeleteDuplicates(t *testing.T) {
	input := []string{"first", "first", "first"}
	expected := "first"

	output := DeleteDuplicates(input, &Options{})

	if len(output) != 1 || output[0] != "first" {
		t.Errorf("Call: DeleteDuplicates(%s). Expected: %s. Actual: %s", fString(input), expected, fString(output))
	}
}

func TestDeleteDuplicatesIgnoreLetterCase(t *testing.T) {
	opt := Options{fFlag: true}
	input := []string{"first", "FIRST", "fiRSt"}
	expected := []string{"first"}

	output := DeleteDuplicates(input, &opt)

	if len(output) != 1 || output[0] != expected[0] {
		t.Errorf("Call: DeleteDuplicates(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
	}
}
