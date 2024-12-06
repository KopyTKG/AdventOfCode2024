package day06

import (
	"aoc2024/aoc"
	"aoc2024/stream"
	"strconv"
)

type Solution struct{}

type direction struct {
	x int
	y int
}

type player struct {
	x int
	y int

	heading direction
}

type headings struct {
	current    int
	directions []direction
}

func (h *headings) next() direction {
	if h.current >= len(h.directions)-1 {
		h.current = 0
	} else {
		h.current++
	}
	return h.directions[h.current]
}

func (p *player) valid(area [][]string) bool {
	x := p.x + p.heading.x
	y := p.y + p.heading.y

	checkY := y > len(area)-1 || y < 0
	checkX := x > len(area[0])-1 || x < 0

	return !checkX && !checkY
}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)
	h := headings{directions: []direction{direction{x: 0, y: -1}, direction{x: 1, y: 0}, direction{x: 0, y: 1}, direction{x: -1, y: 0}}}
	var area [][]string
	var p player
	for y, line := range lines {
		var row []string
		for x, rune := range line {
			if string(rune) == "^" {
				p.x = x
				p.y = y
				p.heading = direction{x: 0, y: -1}
			}
			row = append(row, string(rune))
		}
		area = append(area, row)
	}
	positons := 0
	stop := false
	for !stop {
		if !p.valid(area) {
			positons++
			area[p.y][p.x] = "@"
			stop = true
		} else {
			if area[p.y+p.heading.y][p.x+p.heading.x] != "#" {
				if area[p.y][p.x] != "@" {
					positons++
					area[p.y][p.x] = "@"
				}
				p.x += p.heading.x
				p.y += p.heading.y
			} else {
				p.heading = h.next()
			}
		}
	}
	return strconv.Itoa(positons), nil
}

func (s Solution) Star2(input string) (string, error) {
	return "", aoc.NIP()
}
