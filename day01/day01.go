package day01

import (
	"aoc2024/stream"
	"strconv"
	"strings"
)

type Solution struct{}

func Buble(arr *[]int) {
	sorted := false
	top := len(*arr)
	for !sorted {
		for x := 0; x < top-1; x++ {
			curr := (*arr)[x]
			next := (*arr)[x+1]
			if curr > next {
				(*arr)[x] = next
				(*arr)[x+1] = curr
			}
		}
		if top == 1 {
			sorted = true
		}
		top--
	}
}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)

	var left []int
	var right []int

	for _, bLine := range lines {
		line := string(bLine)
		items := strings.Split(line, "   ")

		lInt, err := strconv.Atoi(items[0])
		if err != nil {
			return "", err
		}
		left = append(left, lInt)

		rInt, err := strconv.Atoi(items[1])
		if err != nil {
			return "", err
		}
		right = append(right, rInt)
	}

	Buble(&left)
	Buble(&right)

	total := 0

	for x := 0; x < len(left); x++ {
		diff := left[x] - right[x]
		if diff < 0 {
			diff *= -1
		}
		total += diff
	}

	return strconv.Itoa(total), nil

}

func (s Solution) Star2(input string) (string, error) {
	lines := stream.ReadLines(input)

	var left []int
	var right []int

	for _, bLine := range lines {
		line := string(bLine)
		items := strings.Split(line, "   ")

		lInt, err := strconv.Atoi(items[0])
		if err != nil {
			return "", err
		}
		left = append(left, lInt)

		rInt, err := strconv.Atoi(items[1])
		if err != nil {
			return "", err
		}
		right = append(right, rInt)
	}

	Buble(&left)
	Buble(&right)

	total := 0

	for _, x := range left {
		count := 0
		for _, y := range right {
			if y == x {
				count++
			}
		}

		total += x * count
	}

	return strconv.Itoa(total), nil
}
