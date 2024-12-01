package day1

import (
	"aoc2024/stream"
	"fmt"
	"strconv"
	"strings"
)

func Run(f string) {
	star1(f)
	star2(f)
}

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

func star1(f string) {
	fmt.Println("Day 1 - Star 1")
	data := stream.ReadBytes(f)
	lines := stream.BtoSa(data)

	var left []int
	var right []int

	for _, bLine := range lines {
		line := string(bLine)
		items := strings.Split(line, "   ")

		lInt, err := strconv.Atoi(items[0])
		if err != nil {
			panic(err)
		}
		left = append(left, lInt)

		rInt, err := strconv.Atoi(items[1])
		if err != nil {
			panic(err)
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

	fmt.Println(total)

}

func star2(f string) {
	fmt.Println("Day 1 - Star 2")
	data := stream.ReadBytes(f)
	lines := stream.BtoSa(data)

	var left []int
	var right []int

	for _, bLine := range lines {
		line := string(bLine)
		items := strings.Split(line, "   ")

		lInt, err := strconv.Atoi(items[0])
		if err != nil {
			panic(err)
		}
		left = append(left, lInt)

		rInt, err := strconv.Atoi(items[1])
		if err != nil {
			panic(err)
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

	fmt.Println(total)
}
