package main

import (
	"fmt"
	"testing"
)

func FString(input []string) string {
	var formatStr string
	for _, str := range input {
		formatStr += fmt.Sprintf("%s, ", str)
	}
	return formatStr
}

func TestParse(t *testing.T) {
	input := "(6+10-4)/(1+1*2)+1"
	expected := []string{"(", "6", "+", "10", "-", "4", ")", "/", "(", "1", "+", "1", "*", "2", ")", "+", "1"}

	output, err := Parse(input)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	for i, str := range expected {
		if str != output[i] {
			t.Errorf("Call: ConvertToRpn(%s). Expected: %s. Actual: %s", input, FString(expected), FString(output))
		}
	}
}

func TestParseWithWhiteSpaces(t *testing.T) {
	input := "(6 + 10 - 4) / (1 + 1 * 2) + 1"
	expected := []string{"(", "6", "+", "10", "-", "4", ")", "/", "(", "1", "+", "1", "*", "2", ")", "+", "1"}

	output, err := Parse(input)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	for i, str := range expected {
		if str != output[i] {
			t.Errorf("Call: ConvertToRpn(%s). Expected: %s. Actual: %s", input, FString(expected), FString(output))
		}
	}
}

func TestParseInvalidString(t *testing.T) {
	input := "   (((123 + 1)"

	_, err := Parse(input)

	if err == nil {
		t.Errorf("Expected errors")
	}
}

func TestConvertToRpn(t *testing.T) {
	input := []string{"(", "6", "+", "10", "-", "4", ")", "/", "(", "1", "+", "1", "*", "2", ")", "+", "1"}
	expected := []string{"6", "10", "+", "4", "-", "1", "1", "2", "*", "+", "/", "1", "+"}

	output := ConvertToRpn(input)

	for i, str := range expected {
		if str != output[i] {
			t.Errorf("Call: ConvertToRpn(%s). Expected: %s. Actual: %s", FString(input), FString(expected), FString(output))
		}
	}
}

func TestConvertToRpnInvalidString(t *testing.T) {
	input := []string{"+qwe1", "2", "  ", "---"}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Unexpected error: %s", r.(error))
		}
	}()

	ConvertToRpn(input)
}

func TestCalculate(t *testing.T) {
	input := []string{"6", "10", "+", "4", "-", "1", "1", "2", "*", "+", "/", "1", "+"}
	expected := 5

	output, err := Calculate(input)

	if expected != output || err != nil {
		t.Errorf("Call: Calculate(%s). Expected: %d. Actual: %d", FString(input), expected, output)
	}
}

func TestCalculateInvalidInput(t *testing.T) {
	input := []string{"+qwe1", "2", "  ", "---"}

	_, err := Calculate(input)

	if err == nil {
		t.Errorf("Expected errors")
	}
}
