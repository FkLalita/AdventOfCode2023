package main

import (
	"strings"
	"testing"
)

func TestCubeCount(t *testing.T) {
	// Test data
	testData := `Game 1: 1 blue; 4 green, 5 blue; 11 red, 3 blue, 11 green; 1 red, 10 green, 4 blue; 17 red, 12 green, 7 blue; 3 blue, 19 green, 15 red
Game 2: 17 red, 10 green; 3 blue, 17 red, 7 green; 10 green, 1 blue, 10 red; 7 green, 15 red, 1 blue; 7 green, 8 blue, 16 red; 18 red, 5 green, 3 blue
Game 3: 10 blue, 3 green, 8 red; 15 green, 14 blue, 1 red; 8 blue, 11 red, 2 green; 5 red, 9 green, 6 blue
Game 4: 1 red, 3 blue; 3 blue, 3 green, 1 red; 11 blue, 2 green; 2 green, 14 blue; 1 green, 7 blue; 11 blue, 5 green
Game 5: 9 green, 5 red, 10 blue; 9 red, 4 blue, 12 green; 9 green, 6 blue, 14 red; 16 green, 8 red, 6 blue; 11 blue, 13 red, 1 green`

	// Call the CubeCount function with the test data
	result := CubeCount(strings.NewReader(testData))

	// Expected sum of IDs for the provided test data
	expectedResult := 0 + 0 + 0 + 0 + 0

	// Check if the result matches the expected value
	if result != expectedResult {
		t.Errorf("Unexpected result. Expected: %d, Got: %d", expectedResult, result)
	}
}
