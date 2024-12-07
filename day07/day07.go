package day07

import (
	"aoc2024/stream"
	"strconv"
	"strings"
)

type Solution struct{}

type calibration struct {
	result      int
	input       []int
	binLen      int
	operators2s [][]rune
	operators3s [][]rune
}

func (c *calibration) Init() {
	if len(c.input) > 0 {
		c.binLen = len(c.input) - 1
		c.buildOperators2d()
		c.buildOperators3d()
	}
}

func (c *calibration) buildOperators2d() {
	c.operators2s = make([][]rune, 1<<c.binLen)
	for i := range c.operators2s {
		c.operators2s[i] = make([]rune, c.binLen)
	}

	var generate func(index, digit int)
	generate = func(index, digit int) {
		if digit == c.binLen {
			return
		}
		half := 1 << (c.binLen - digit - 1)
		for i := 0; i < 1<<c.binLen; i++ {
			if i/half%2 == 0 {
				c.operators2s[i][digit] = '+'
			} else {
				c.operators2s[i][digit] = '*'
			}
		}
		generate(index, digit+1)
	}

	generate(0, 0)
}

func (c *calibration) buildOperators3d() {
	c.operators3s = make([][]rune, 1<<(2*c.binLen))
	for i := range c.operators3s {
		c.operators3s[i] = make([]rune, c.binLen)
	}

	var generate func(index, digit int)
	generate = func(index, digit int) {
		if digit == c.binLen {
			return
		}
		third := 1 << (2 * (c.binLen - digit - 1))
		for i := 0; i < 1<<(2*c.binLen); i++ {
			switch (i / third) % 3 {
			case 0:
				c.operators3s[i][digit] = '+'
			case 1:
				c.operators3s[i][digit] = '*'
			case 2:
				c.operators3s[i][digit] = '|'
			}
		}
		generate(index, digit+1)
	}

	generate(0, 0)
}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)

	var calibrations []calibration

	for _, line := range lines {
		sum, err := strconv.Atoi(strings.Split(line, ":")[0])
		if err != nil {
			return "", err
		}

		sItems := strings.Split(strings.Split(line, ":")[1], " ")
		var items []int
		for i, sItem := range sItems {
			if i == 0 {
				continue
			}
			item, err := strconv.Atoi(sItem)
			if err != nil {
				return "", err
			}
			items = append(items, item)

		}
		tmp := calibration{result: sum, input: items}
		tmp.Init()
		calibrations = append(calibrations, tmp)
	}
	total := 0
	for _, cal := range calibrations {
		for k := 0; k < len(cal.operators2s); k++ {
			res := cal.input[0]
			for x := 0; x < cal.binLen; x++ {
				if cal.operators2s[k][x] == '+' {
					res += cal.input[x+1]
				} else {
					res *= cal.input[x+1]
				}
			}
			if res == cal.result {
				total += cal.result
				break
			}
		}
	}
	return strconv.Itoa(total), nil
}

func (s Solution) Star2(input string) (string, error) {
	lines := stream.ReadLines(input)

	var calibrations []calibration

	for _, line := range lines {
		sum, err := strconv.Atoi(strings.Split(line, ":")[0])
		if err != nil {
			return "", err
		}

		sItems := strings.Split(strings.Split(line, ":")[1], " ")
		var items []int
		for i, sItem := range sItems {
			if i == 0 {
				continue
			}
			item, err := strconv.Atoi(sItem)
			if err != nil {
				return "", err
			}
			items = append(items, item)

		}
		tmp := calibration{result: sum, input: items}
		tmp.Init()
		calibrations = append(calibrations, tmp)
	}
	total := 0
	for _, cal := range calibrations {
		for k := 0; k < len(cal.operators3s); k++ {
			res := cal.input[0]
			for x := 0; x < len(cal.operators3s[k]); x++ {

				switch cal.operators3s[k][x] {
				case '+':
					res += cal.input[x+1]
				case '*':
					res *= cal.input[x+1]
				case '|':
					i, err := strconv.Atoi(strconv.Itoa(res) + strconv.Itoa(cal.input[x+1]))
					if err != nil {
						return "", err
					}
					res = i

				}

			}
			if res == cal.result {
				total += cal.result
				break
			}
		}
	}
	return strconv.Itoa(total), nil
}
