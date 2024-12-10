package day09

import (
	"aoc2024/stream"
	"strconv"
)

type Solution struct{}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)

	var disk []int
	idCounter := 0
	for _, line := range lines {
		for i, char := range line {
			if i%2 == 0 {
				s, err := strconv.Atoi(string(char))
				if err != nil {
					return "", err
				}
				for range s {
					disk = append(disk, idCounter)
				}
				idCounter++
			} else {
				s, err := strconv.Atoi(string(char))
				if err != nil {
					return "", err
				}
				for range s {
					disk = append(disk, -1)
				}
			}
		}
	}
	for x := 0; x < len(disk); x++ {
		if disk[x] == -1 {
			for k := len(disk) - 1; k > x; k-- {
				if disk[k] != -1 {
					disk[x] = disk[k]
					disk[k] = -1
					break
				}
			}
		}
	}
	sum := 0
	for i, f := range disk {
		if f == -1 {
			break
		}
		sum += i * f
	}

	return strconv.Itoa(sum), nil
}

func reorganizeDisk(disk [][]int) [][]int {
	for x := 0; x < len(disk); x++ {
		if len(disk[x]) > 0 && disk[x][0] == -1 {
			for k := len(disk) - 1; k > x; k-- {
				if len(disk[k]) > 0 && disk[k][0] != -1 {
					if len(disk[k]) <= len(disk[x]) {
						swapOrReorganizeSlices(disk, x, k)
						break
					}
				}
			}
		}
	}
	return disk
}

func swapOrReorganizeSlices(disk [][]int, x, k int) {
	diff := len(disk[x]) - len(disk[k])
	if diff == 0 {
		disk[x], disk[k] = disk[k], disk[x]
	} else {
		size := len(disk[x]) - diff
		disk[x] = disk[k]

		emptySlice := make([]int, diff)
		for i := range emptySlice {
			emptySlice[i] = -1
		}

		disk[k] = make([]int, size)
		for i := range disk[k] {
			disk[k][i] = -1
		}
		disk = append(disk[:x+1], append([][]int{emptySlice}, disk[x+1:]...)...)
	}
}

func (s Solution) Star2(input string) (string, error) {
	lines := stream.ReadLines(input)

	var disk [][]int
	idCounter := 0
	for _, line := range lines {
		for i, char := range line {
			if i%2 == 0 {
				s, err := strconv.Atoi(string(char))
				if err != nil {
					return "", err
				}
				var tmp []int
				for range s {
					tmp = append(tmp, idCounter)
				}
				disk = append(disk, tmp)

				idCounter++
			} else {
				s, err := strconv.Atoi(string(char))
				if err != nil {
					return "", err
				}
				var tmp []int
				for range s {
					tmp = append(tmp, -1)
				}
				disk = append(disk, tmp)
			}
		}
	}

	reorganizeDisk(disk)

	var clean []int
	for _, f := range disk {
		for _, a := range f {
			clean = append(clean, a)
		}
	}
	sum := 0
	for i, f := range clean {
		if f == -1 {
			continue
		}

		sum += i * f
	}

	return strconv.Itoa(sum), nil
}
