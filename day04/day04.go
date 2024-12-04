package day04

import (
	"aoc2024/stream"
	"strconv"
)

type Solution struct{}

type word struct {
	start     []int
	direction []int
}

type search struct {
	// public
	Grid  [][]string
	Found []word
	Word  []string

	// private
	cols       int
	rows       int
	directions [][]int
}

func (s *search) Init(input string) {
	tmp := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	s.directions = tmp
	s.cols = 0
	s.rows = 0

	for _, rune := range input {
		(*s).Word = append((*s).Word, string(rune))
	}
}

func (s *search) CalculateGrid() {
	if len(s.Grid) > 0 {
		s.rows = len(s.Grid)
		s.cols = len(s.Grid[0])
	}
}

func (s *search) searchInDirection(row, col, dirRow, dirCol int) bool {
	for i := 0; i < len(s.Word); i++ {
		newRow, newCol := row+i*dirRow, col+i*dirCol
		if newRow < 0 || newRow >= s.rows || newCol < 0 || newCol >= s.cols || s.Grid[newRow][newCol] != s.Word[i] {
			return false
		}
	}
	return true
}

func (s *search) searchXDirection(row, col int, dir [][]int) bool {
	i := 0
	for _, coords := range dir {
		newRow, newCol := row+coords[0], col+coords[1]
		if newRow < 0 || newRow >= s.rows || newCol < 0 || newCol >= s.cols || s.Grid[newRow][newCol] != s.Word[i] {
			return false
		}
		i += 2
	}
	return true
}

func (s *search) FindWord() {
	for row := 0; row < s.rows; row++ {
		for col := 0; col < s.cols; col++ {
			for _, dir := range s.directions {
				if s.searchInDirection(row, col, dir[0], dir[1]) {
					tmp := word{start: []int{row, col}, direction: dir}
					s.Found = append(s.Found, tmp)
				}
			}
		}
	}
}

func reverse(i [][]int) [][]int {
	var tmp [][]int
	for x := len(i) - 1; x > -1; x-- {
		tmp = append(tmp, (i)[x])
	}
	return tmp
}

func (s *search) FindX() {
	tLbR := [][]int{{-1, -1}, {1, 1}}
	bLtR := [][]int{{1, -1}, {-1, 1}}

	for row := 0; row < s.rows; row++ {
		for col := 0; col < s.cols; col++ {
			isX := true
			if s.Grid[row][col] == s.Word[len(s.Word)-2] {
				if !s.searchXDirection(row, col, tLbR) && !s.searchXDirection(row, col, reverse(tLbR)) {
					isX = false
				}
				if !s.searchXDirection(row, col, bLtR) && !s.searchXDirection(row, col, reverse(bLtR)) {
					isX = false
				}
				if isX {
					tmp := word{start: []int{row, col}, direction: []int{row, col}}
					s.Found = append(s.Found, tmp)
				}
			}
		}
	}
}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)
	data := new(search)
	data.Init("XMAS")

	for y := 0; y < len(lines); y++ {
		line := lines[y]
		var row []string
		for x := 0; x < len(line); x++ {
			current := string(line[x])
			row = append(row, current)
		}
		data.Grid = append(data.Grid, row)
	}

	data.CalculateGrid()
	data.FindWord()

	return strconv.Itoa(len(data.Found)), nil
}

func (s Solution) Star2(input string) (string, error) {
	lines := stream.ReadLines(input)
	data := new(search)
	data.Init("MAS")

	for y := 0; y < len(lines); y++ {
		line := lines[y]
		var row []string
		for x := 0; x < len(line); x++ {
			current := string(line[x])
			row = append(row, current)
		}
		data.Grid = append(data.Grid, row)
	}

	data.CalculateGrid()
	data.FindX()
	return strconv.Itoa(len(data.Found)), nil
}
