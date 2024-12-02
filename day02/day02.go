package day02

import (
	"aoc2024/stream"
	"strconv"
	"strings"
)

type Solution struct{}

func Pop(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func toAsc(arr *[]int) {
	if (*arr)[0] < (*arr)[1] {
		return
	}
	var tmp []int

	for x := len(*arr) - 1; x > -1; x-- {
		tmp = append(tmp, (*arr)[x])
	}
	for i := range tmp {
		(*arr)[i] = tmp[i]
	}
}

func copyArr(src []int) []int {
	var tmp []int

	for _, i := range src {
		tmp = append(tmp, i)
	}
	return tmp
}

func isValidSequence(numbers []int, damper bool) bool {
	toAsc(&numbers)

	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		diffcheck := (1 <= diff && diff <= 3)
		if !diffcheck {
			if !damper {
				return false
			}

			for x := range numbers {
				if isValidSequence(Pop(copyArr(numbers), x), false) {
					return true
				}
			}
			return false
		}
	}

	return true
}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)
	var numbers [][]int

	for _, line := range lines {
		items := strings.Split(strings.TrimSpace(line), " ")
		var row []int

		for _, item := range items {
			num, err := strconv.Atoi(item)
			if err != nil {
				return "", err
			}

			row = append(row, num)
		}

		numbers = append(numbers, row)
	}

	count := 0

	for _, row := range numbers {
		if isValidSequence(row, false) {
			count++
		}
	}

	return strconv.Itoa(count), nil
}

func (s Solution) Star2(input string) (string, error) {
	lines := stream.ReadLines(input)
	var numbers [][]int

	for _, line := range lines {
		items := strings.Split(strings.TrimSpace(line), " ")
		var row []int

		for _, item := range items {
			num, err := strconv.Atoi(item)
			if err != nil {
				return "", err
			}

			row = append(row, num)
		}

		numbers = append(numbers, row)
	}

	count := 0

	for _, row := range numbers {
		if isValidSequence(row, true) {
			count++
		}
	}
	return strconv.Itoa(count), nil
}
