package day03

import (
	"aoc2024/aoc"
	"aoc2024/stream"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Solution struct{}

type smul struct {
	start int
	code  string
}

type smuls []smul

func (s *smuls) pop(i int) {
	(*s) = append((*s)[:i], (*s)[i+1:]...)
}

func (s *smuls) remove(start int, end int) {
	(*s) = append((*s)[:start], (*s)[end:]...)
}

func (s *smuls) getFirst(item string) (int, error) {
	for i, slice := range *s {
		if slice.code == item {
			return i, nil
		}
	}
	return -1, errors.New("Out of bounds")
}

type mul struct {
	a   int
	b   int
	res int
}

func (s *smuls) sort() {
	sorted := false
	top := len(*s)
	for !sorted {
		for x := 0; x < top-1; x++ {
			curr := (*s)[x]
			next := (*s)[x+1]
			if curr.start > next.start {
				(*s)[x] = next
				(*s)[x+1] = curr
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
	matcher := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	dos := regexp.MustCompile(`do()`)
	donts := regexp.MustCompile(`don't()`)

	var matches []string

	for _, line := range lines {
		var cmds smuls

		idos := dos.FindAllStringIndex(line, -1)
		idonts := donts.FindAllStringIndex(line, -1)
		for _, do := range idos {
			fmt.Println(do)
			tmp := smul{start: do[0], code: "do"}
			cmds = append(cmds, tmp)
		}
		for _, dont := range idonts {
			tmp := smul{start: dont[0], code: "dont"}
			cmds = append(cmds, tmp)
		}
		cmds.sort()

		found := matcher.FindAllString(line, -1)
		indexes := matcher.FindAllStringIndex(line, -1)

		write := true
		for x := 0; x < len(indexes); x++ {
			position := indexes[x][0]
			doindex, err := cmds.getFirst("do")
			if err != nil {
				break
			}
			dontindex, err := cmds.getFirst("dont")
			if err != nil {
				break
			}
			if position < cmds[dontindex].start && !write {
				if position > cmds[doindex].start {
					write = true
				}
			}

			if write {
				tmp := found[x]
				matches = append(matches, tmp)
			}
			fmt.Println(cmds)
			cmds.pop(doindex)
			cmds.pop(dontindex)

		}

	}

	var multiplications []mul
	fmt.Println(len(matches))
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

	//return strconv.Itoa(total), nil
	return "", aoc.NIP()
}
