package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ExtractDigitsAndSpelledStrings extracts digits from a string and converts spelled-out numbers to digits.
func ExtractDigitsAndSpelledStrings(str string) []int {
	spellingMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	digits := []int{}
	for i, char := range str {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			digits = append(digits, digit)
		} else {
			for spelling, number := range spellingMap {
				if strings.HasPrefix(str[i:], spelling) {
					digits = append(digits, number)
					break
				}
			}
		}
	}

	return digits
}

// GetSum calculates the sum of all the results once the digits are extracted.
func GetSum(lines []string, extractFunc func(string) []int) int {
	answer := 0
	for _, line := range lines {
		digits := extractFunc(line)
		if len(digits) > 0 {
			first, last := digits[0], digits[len(digits)-1]
			answer += first*10 + last
		} else {
			fmt.Println("No digits found in the string")
		}
	}
	return answer
}

// ExtractDigits extracts only digits from a string.
func ExtractDigits(str string) []int {
	digits := []int{}
	for _, char := range str {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			digits = append(digits, digit)
		}
	}
	return digits
}

func main() {
	filename := "input.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	// First Challenge
	fmt.Println("First Challenge Answer:", GetSum(lines, ExtractDigits))
	// Second Challenge
	fmt.Println("Second Challenge Answer:", GetSum(lines, ExtractDigitsAndSpelledStrings))
}
