package aoc

import (
	"errors"
	"fmt"
	"os"
)

func NIP() error {
	return errors.New("Not implemented")
}

type DailySolution interface {
	Star1(input string) (string, error)
	Star2(input string) (string, error)
}

func Run(day int, solution DailySolution) {
	files := []string{"test.txt", "input.txt"}

	for _, file := range files {
		input := fmt.Sprintf("day%02d/%s", day, file)
		_, err := os.Stat(input)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", file, err)
			continue
		}

		fmt.Printf("--- Day %d (%s) ---\n", day, file)

		result1, err := solution.Star1(input)
		if err != nil {
			fmt.Printf("Star 1 Error: %v\n", err)
		} else {
			fmt.Printf("Star 1: %s\n", result1)
		}

		result2, err := solution.Star2(input)
		if err != nil {
			fmt.Printf("Star 2 Error: %v\n", err)
		} else {
			fmt.Printf("Star 2: %s\n", result2)
		}

		fmt.Println()
	}
}
