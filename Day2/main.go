package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	CubeCount()
}

func CubeCount() {
	redCount := 12
	greenCount := 13
	blueCount := 14

	inputs, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error Reading file:", err)
		return
	}
	input := string(inputs)

	games := strings.Split(input, "Game ")

	// Variable to store the sum of possible game IDs
	sumOfPossibleIDs := 0

	// loop through each game
	for _, gameInfo := range games[1:] {
		// Split game information into subsets
		subsets := strings.Split(gameInfo, ":")

		// Loop through each subset
		for _, subset := range subsets {
			// Extract color and count from the subset
			parts := strings.Fields(subset)
			if len(parts) == 2 {
				color := parts[1]
				countStr := parts[0]

				// Convert count string to integer
				count, err := strconv.Atoi(countStr)
				if err != nil {
					fmt.Println("Error converting string:", err)
					return
				}
				// deduct the count from the corresponding cube count
				switch color {
				case "red":
					redCount -= count
				case "green":
					greenCount -= count
				case "blue":
					blueCount -= count
				default:
					fmt.Println("Unknown color:", color)
				}
			}

		}
		fmt.Println("red", redCount, blueCount, greenCount)

		// check if the remaining counts are non-negative/positive
		if redCount >= 0 && greenCount >= 0 && blueCount >= 0 {
			// if yes, the game is possible with the given cube counts
			// add the game's id to the sum
			gameID := strings.Split(gameInfo, ":")[0]
			// Convert the game ID to an integer
			id, err := strconv.Atoi(gameID)
			if err != nil {
				fmt.Println("Error converting game ID to integer:", err)
				return
			}
			sumOfPossibleIDs += id
			//fmt.Println("possible game id:", gameID)
		}

		// reset cube counts for the next game
		redCount = 12
		greenCount = 13
		blueCount = 14
	}

	// Print the sum of all possible game IDs
	fmt.Println("Sum of Possible Game IDs:", sumOfPossibleIDs)
}
