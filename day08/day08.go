package day08

import (
	"aoc2024/stream"
	"math"
	"strconv"
)

type Solution struct{}

type point struct {
	x int
	y int
}

type antenna struct {
	code  string
	cords point
}

type antinode struct {
	cords point
}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)

	points := make(map[point]antenna)
	antinodes := make(map[antinode]bool)
	minX, maxX, minY, maxY := 0, len(lines[0]), 0, len(lines)

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			char := string(lines[y][x])
			if char != "." {
				points[point{x: x, y: y}] = antenna{code: char, cords: point{x: x, y: y}}
			}
		}
	}
	for _, ant := range points {
		for y := minY; y < maxY; y++ {
			for x := minX; x < maxX; x++ {
				selected := points[point{x, y}]

				if selected == ant {
					continue
				}

				if selected.code == ant.code {
					diffX, diffY := ant.cords.x-selected.cords.x, ant.cords.y-selected.cords.y

					diffX = int(math.Abs(float64(diffX)))
					diffY = int(math.Abs(float64(diffY)))

					ax1, ax2, ay1, ay2 := 0, 0, 0, 0

					if selected.cords.x > ant.cords.x {
						ax1, ax2 = ant.cords.x-diffX, selected.cords.x+diffX
					} else {

						ax1, ax2 = ant.cords.x+diffX, selected.cords.x-diffX
					}

					if selected.cords.y > ant.cords.y {
						ay1, ay2 = ant.cords.y-diffY, selected.cords.y+diffY
					} else {

						ay1, ay2 = ant.cords.y+diffY, selected.cords.y-diffY
					}

					if (ax1 >= minX && ax1 < maxX) && (ay1 >= minY && ay1 < maxY) {
						if !antinodes[antinode{point{ax1, ay1}}] {
							antinodes[antinode{point{ax1, ay1}}] = true
						}
					}

					if (ax2 >= minX && ax2 < maxX) && (ay2 >= minY && ay2 < maxY) {
						if !antinodes[antinode{point{ax2, ay2}}] {
							antinodes[antinode{point{ax2, ay2}}] = true
						}
					}
				}
			}
		}
	}

	return strconv.Itoa(len(antinodes)), nil
}

func (s Solution) Star2(input string) (string, error) {
	lines := stream.ReadLines(input)
	minX, maxX, minY, maxY := 0, len(lines[0]), 0, len(lines)

	// Group antennas by frequency
	antennasByFreq := make(map[string][]point)
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			char := string(lines[y][x])
			if char != "." {
				antennasByFreq[char] = append(antennasByFreq[char], point{x: x, y: y})
			}
		}
	}

	antinodes := make(map[point]bool)

	// Calculate antinodes for each frequency group
	for _, antennas := range antennasByFreq {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				a1, a2 := antennas[i], antennas[j]
				dx, dy := a2.x-a1.x, a2.y-a1.y

				// Check antinodes in both directions
				for k := -50; k <= 50; k++ {
					ax := a1.x + dx*k
					ay := a1.y + dy*k
					if ax < minX || ax >= maxX || ay < minY || ay >= maxY {
						continue
					}
					antinodes[point{x: ax, y: ay}] = true
				}
			}
		}
	}

	return strconv.Itoa(len(antinodes)), nil
}
