package day10

import (
	"aoc2024/aoc"
	"aoc2024/stream"
	"aoc2024/vectors"
	"fmt"
	"strconv"
)

type Solution struct{}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)

	var board [][]int
	for _, line := range lines {
		var row []int
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return "", err
			}
			row = append(row, num)
		}
		board = append(board, row)
	}

	fmt.Println(board)

	var steps []vectors.Vector2d
	return "", aoc.NIP()
}

func (s Solution) Star2(input string) (string, error) {
	return "", aoc.NIP()
}
