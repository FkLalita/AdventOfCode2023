package main

import (
	"bufio"
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
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error Reading file:", err)
		return
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)
	sumOfPossibleIDs := 0

	for scanner.Scan() {
		line := scanner.Text()
		gameID := strings.Split(line, ":")
		game_ids := gameID[0]
		game_id, err := strconv.Atoi(game_ids[5:])
		if err != nil {
			fmt.Println(err)
		}

		var gameSatisfiesConditions bool

		for _, gameIDs := range gameID[1:] {
			subset := strings.Split(gameIDs, ";")
			for _, subsets := range subset {
				games := strings.Split(subsets, ",")

				for _, colorz := range games {
					colors := strings.Fields(colorz)
					color := colors[1]
					count, _ := strconv.Atoi(colors[0])
					switch {
					case color == "blue" && count > 14:
						gameSatisfiesConditions = false
					case color == "red" && count > 12:
						gameSatisfiesConditions = false
					case color == "green" && count > 13:
						gameSatisfiesConditions = false
					default:
						gameSatisfiesConditions = true
					}

					if !gameSatisfiesConditions {
						break
					}
				}

				if !gameSatisfiesConditions {
					break
				}
			}

			if !gameSatisfiesConditions {
				break
			}
		}

		if gameSatisfiesConditions {
			sumOfPossibleIDs += game_id
		}
	}
	fmt.Println(sumOfPossibleIDs)
}
