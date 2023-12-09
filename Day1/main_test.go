package main

import (
	"strings"
	"testing"
)

// TestExtractDigitsAndSpelledStrings tests the ExtractDigitsAndSpelledStrings function.
func TestExtractDigitsAndSpelledStrings(t *testing.T) {
	// Test case where the input contains both digits and spelled-out numbers
	input := "abc123def four ghi56seven"
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	result := ExtractDigitsAndSpelledStrings(input)

	if len(result) != len(expected) {
		t.Errorf("Length of result and expected do not match.")
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Mismatch at index %d. Expected: %d, Got: %d", i, expected[i], result[i])
		}
	}
}

// TestExtractDigits tests the ExtractDigits function.
func TestExtractDigits(t *testing.T) {
	// Test case where the input contains only digits
	input := "abc123def456ghi789"
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := ExtractDigits(input)

	if len(result) != len(expected) {
		t.Errorf("Length of result and expected do not match.")
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Mismatch at index %d. Expected: %d, Got: %d", i, expected[i], result[i])
		}
	}
}

// TestGetSum tests the GetSum function.
func TestGetSum(t *testing.T) {
	lines := []string{"abc123def", "fourghi56seven", "xyz789uvw"}
	expected := 123 + 7 + 9 // Sum of the first digits and last digits in each line

	// Test with ExtractDigitsAndSpelledStrings
	result := GetSum(lines, ExtractDigitsAndSpelledStrings)
	if result != expected {
		t.Errorf("Expected: %d, Got: %d", expected, result)
	}

	// Test with ExtractDigits
	result = GetSum(lines, ExtractDigits)
	if result != expected {
		t.Errorf("Expected: %d, Got: %d", expected, result)
	}
}

func TestMain(t *testing.T) {
	// Set up test input data
	input := "abc123def four ghi56seven\nfourghi56sevenxyz789uvw"
	lines := strings.Split(input, "\n")

	// Redirect stdout for testing output
	output := captureOutput(func() {
		// Test main function
		main()
	})

	// Test expected output
	expectedOutput := "First Challenge Answer: 27\nSecond Challenge Answer: 27\n"
	if output != expectedOutput {
		t.Errorf("Expected Output:\n%s\nGot Output:\n%s", expectedOutput, output)
	}
}

