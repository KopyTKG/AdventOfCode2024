package day03

import (
	"aoc2024/stream"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Solution struct{}

type mul struct {
	a   int
	b   int
	res int
}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)
	matcher := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	var matches []string
	for _, line := range lines {
		found := matcher.FindAllString(line, -1)
		matches = append(matches, found...)
	}

	var multiplications []mul
	for _, match := range matches {
		cleaned := strings.ReplaceAll(match, "mul(", ")")
		cleaned = strings.ReplaceAll(cleaned, ")", "")

		parts := strings.Split(cleaned, ",")

		if len(parts) < 2 {
			return "", errors.New("Missing args")
		}

		a, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", err
		}

		b, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}

		tmp := mul{a: a, b: b, res: a * b}
		multiplications = append(multiplications, tmp)
	}

	total := 0

	for _, item := range multiplications {
		total += item.res
	}

	return strconv.Itoa(total), nil
}

func (s Solution) Star2(input string) (string, error) {
	lines := stream.ReadLines(input)
	matcher := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`)
	var matches []string
	write := true
	for _, line := range lines {

		found := matcher.FindAllString(line, -1)

		for _, match := range found {
			if match == "don't()" {
				write = false
			}

			if match == "do()" {
				write = true
			}

			if write && string(match[0]) == "m" {
				tmp := match
				matches = append(matches, tmp)
			}

		}

	}

	var multiplications []mul
	for _, match := range matches {
		cleaned := strings.ReplaceAll(match, "mul(", ")")
		cleaned = strings.ReplaceAll(cleaned, ")", "")

		parts := strings.Split(cleaned, ",")

		if len(parts) < 2 {
			return "", errors.New("Missing args")
		}

		a, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", err
		}

		b, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}

		tmp := mul{a: a, b: b, res: a * b}
		multiplications = append(multiplications, tmp)
	}

	total := 0

	for _, item := range multiplications {
		total += item.res
	}

	return strconv.Itoa(total), nil
}
