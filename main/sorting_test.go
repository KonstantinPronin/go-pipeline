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
	*f = true
	*u = true
	*r = true
	*k = 1
	defer func() {
		*f = false
		*u = false
		*r = false
		*k = -1
	}()
	input := []string{"hhh bbb", "bbb aaa", "zzz AAA", "ooo yyy", "BBB AAA"}
	expected := []string{"ooo yyy", "hhh bbb", "zzz AAA", "bbb aaa"}

	output, err := Sorting(input)

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
	input := []string{"hhh bbb", "bbb aaa", "zzz AAA"}
	expected := []string{"bbb aaa", "hhh bbb", "zzz AAA"}

	output, _ := StringSorting(input, -1)
	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: StringSorting(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestStringSortingInvalidColumn(t *testing.T) {
	input := []string{"hhh bbb", "bbb aaa", "zzz AAA"}

	_, err := StringSorting(input, 10)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestStringSortingByColumn(t *testing.T) {
	input := []string{"hhh bbb", "bbb aaa", "zzz AAA"}
	expected := []string{"zzz AAA", "bbb aaa", "hhh bbb"}

	output, _ := StringSorting(input, 1)
	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: StringSorting(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestStringSortingByColumnIgnoreLetterCase(t *testing.T) {
	*f = true
	defer func() { *f = false }()
	input := []string{"hhh bbb", "zzz AAA", "bbb aaa"}
	expected := []string{"zzz AAA", "bbb aaa", "hhh bbb"}

	output, _ := StringSorting(input, 1)
	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: StringSorting(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestStringSortingIgnoreLetterCase(t *testing.T) {
	*f = true
	defer func() { *f = false }()
	input := []string{"abc", "ABC"}
	expected := []string{"abc", "ABC"}

	output, _ := StringSorting(input, -1)
	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: StringSorting(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestIntSortingWithError(t *testing.T) {
	input := []string{"abc", "qwe"}
	_, err := IntSorting(input)

	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestIntSortingNoError(t *testing.T) {
	input := []string{"3", "2", "1"}
	expected := []string{"1", "2", "3"}
	output, err := IntSorting(input)

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
	input := []string{"1", "2", "3"}
	expected := []string{"3", "2", "1"}
	output := Reverse(input)

	for i, str := range output {
		if str != expected[i] {
			t.Errorf("Call: Revers(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestDeleteDuplicates(t *testing.T) {
	input := []string{"first", "first", "first"}
	expected := "first"
	output := DeleteDuplicates(input)

	if len(output) != 1 || output[0] != "first" {
		t.Errorf("Call: DeleteDuplicates(%s). Expected: %s. Actual: %s", fString(input), expected, fString(output))
	}
}

func TestDeleteDuplicatesIgnoreLetterCase(t *testing.T) {
	*f = true
	input := []string{"first", "FIRST", "fiRSt"}
	expected := []string{"first"}
	output := DeleteDuplicates(input)

	if len(output) != 1 || output[0] != expected[0] {
		t.Errorf("Call: DeleteDuplicates(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
	}
}
