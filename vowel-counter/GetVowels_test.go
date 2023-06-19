package main

import (
	"testing"
)

func TestGetVowels(t *testing.T) {
	input := "Hello World"
	expectedOutput := "eo"

	output := GetVowels(input)

	if output != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, output)
	}
}

// Add more test cases as needed

func TestMain(t *testing.T) {
	// You can also write tests for the main function if needed
	// However, keep in mind that running tests for the main function
	// may execute the code that relies on user input or external dependencies.
	// In such cases, it's preferable to focus on testing individual functions.
}

