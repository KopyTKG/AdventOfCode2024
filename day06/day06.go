package day06

import (
	"aoc2024/stream"
	"strconv"
)

type Solution struct{}

type direction struct {
	x int
	y int
}

type point struct {
	x        int
	y        int
	headings direction
}
type player struct {
	x int
	y int

	heading direction

	history []point
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

func (p *player) step() {
	p.history = append(p.history, point{x: p.x, y: p.y, headings: p.heading})
	p.x += p.heading.x
	p.y += p.heading.y
}

func (p *player) checkCycle() bool {
	visited := make(map[point]bool)

	for _, x := range p.history {
		if visited[x] {
			return true
		}
		visited[x] = true
	}
	return false
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
				p.step()
			} else {
				p.heading = h.next()
			}
		}
	}
	return strconv.Itoa(positons), nil
}

func runThread(start, stop int, core [][]string, base player, h headings) int {
	positions := 0
	for y := start; y < stop; y++ {
		for x := 0; x < len(core[0]); x++ {
			if core[y][x] == "#" {
				continue
			}

			var area [][]string
			var p player
			for _, line := range core {
				var row []string
				for _, rune := range line {
					row = append(row, rune)
				}
				area = append(area, row)
			}
			area[y][x] = "#"
			p.x = base.x
			p.y = base.y
			p.heading = h.directions[0]
			h.current = 0

			stop := false
			steps := 0
			for !stop {
				if steps == 100 {
					if p.checkCycle() {
						positions++
						break
					}
					steps = 0
				}
				if !p.valid(area) {
					stop = true
				} else {
					if area[p.y+p.heading.y][p.x+p.heading.x] != "#" {
						p.step()
						steps++
					} else {
						p.heading = h.next()
					}
				}
			}

		}
	}
	return positions
}

func (s Solution) Star2(input string) (string, error) {
	lines := stream.ReadLines(input)
	h := headings{directions: []direction{{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}}}
	var core [][]string
	var base player
	for y, line := range lines {
		var row []string
		for x, rune := range line {
			if string(rune) == "^" {
				base.x = x
				base.y = y
				base.heading = direction{x: 0, y: -1}
				row = append(row, ".")
				continue
			}
			row = append(row, string(rune))
		}
		core = append(core, row)
	}

	totalPositions := 0
	numWorkers := 21
	rangeSize := len(core)
	if rangeSize < 20 {
		totalPositions = runThread(0, len(core), core, base, h)
	} else {
		if len(core) > 120 {
			rangeSize = 120 / numWorkers
		}
		results := make(chan int, numWorkers)
		for i := 0; i < numWorkers; i++ {
			start := i * rangeSize
			end := start + rangeSize
			if i == numWorkers-1 {
				end = len(core)
			}
			go func(id, s, e int) {
				result := runThread(s, e, core, base, h)
				results <- result
			}(i, start, end)
		}

		for i := 0; i < numWorkers; i++ {
			threadResult := <-results
			totalPositions += threadResult
		}
	}

	return strconv.Itoa(totalPositions), nil
}
