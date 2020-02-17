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

func TestConvertToRpn(t *testing.T) {
	input := []string{"(", "6", "+", "10", "-", "4", ")", "/", "(", "1", "+", "1", "*", "2", ")", "+", "1"}
	expected := []string{"6", "10", "+", "4", "-", "1", "1", "2", "*", "+", "/", "1", "+"}

	output := ConvertToRpn(input)

	for i, str := range expected {
		if str != output[i] {
			t.Errorf("Call: ConvertToRpn(%s). Expected: %s. Actual: %s", fString(input), fString(expected), fString(output))
		}
	}
}

func TestCalculate(t *testing.T) {
	input := []string{"6", "10", "+", "4", "-", "1", "1", "2", "*", "+", "/", "1", "+"}
	expected := 5

	output, err := Calculate(input)

	if expected != output || err != nil {
		t.Errorf("Call: Calculate(%s). Expected: %d. Actual: %d", fString(input), expected, output)
	}
}
