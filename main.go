package main

import (
	"aoc2024/aoc"
	"aoc2024/day01"
	"aoc2024/day02"
)

func main() {
	aoc.Run(1, day01.Solution{})
	aoc.Run(2, day02.Solution{})
}

/* package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func RemoveIdx(s []int, i int) []int {
	r := make([]int, 0)
	r = append(r, s[:i]...)
	return append(r, s[i+1:]...)
}

func IsSafeSequence(nums []int) bool {
	cmp := func(a, b int) bool { return a < b }
	for i := range nums {
		if nums[i]-nums[i+1] == 0 {
			continue
		}
		if nums[i]-nums[i+1] > 0 {
			cmp = func(a, b int) bool { return a > b }
		}
		break
	}

	for i := 1; i < len(nums); i++ {
		diff := math.Abs(float64(nums[i] - nums[i-1]))

		if diff < 1 || diff > 3 || !cmp(nums[i-1], nums[i]) {
			return false
		}
	}

	return true
}

func IsSafeSequenceDampened(nums []int) bool {
	if IsSafeSequence(nums) {
		return true
	}
	for i, _ := range nums {
		if IsSafeSequence(RemoveIdx(nums, i)) {
			return true
		}
	}
	return false
}

func Day02() (int64, int64) {
	in, err := os.ReadFile("day02/input.txt")
	if err != nil {
		fmt.Print(err)
		return -1, -1
	}

	var safe int64 = 0
	var safeDampened int64 = 0
	for _, line := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		nums := make([]int, 0)
		for _, num := range strings.Fields(line) {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}

		if IsSafeSequence(nums) {
			safe++
		}

		if IsSafeSequenceDampened(nums) {
			safeDampened++
		}
	}

	return safe, safeDampened
}

func main() {
	a, b := Day02()

	fmt.Println(a)
	fmt.Println(b)
} */
