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
