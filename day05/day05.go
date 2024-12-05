package day05

import (
	"aoc2024/stream"
	"strconv"
	"strings"
)

type Solution struct{}

type page struct {
	parent *page
	number int
	child  *page
}

func (p *page) setChild(n *page) {
	if (*p).child == nil {
		(*p).child = n
		n.setParent(p)
	} else {
		(*p).child.setChild(n)
	}
}

func (p *page) setParent(n *page) {
	(*p).parent = n
}

func (p *page) getSteps() []page {
	var steps []page
	current := p
	for {
		steps = append(steps, *current)
		if current.child == nil {
			break
		}
		current = current.child
	}
	return steps
}

type update struct {
	root page
}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)

	rules := make(map[int][]int)
	var updates []update
	var valid []update

	for _, line := range lines {
		r := false
		u := false
		for _, rune := range line {
			if string(rune) == "|" {
				r = true
				break
			} else if string(rune) == "," {
				u = true
				break
			}
		}

		if r {
			pages := strings.Split(line, "|")

			what, err := strconv.Atoi(pages[0])
			if err != nil {
				return "", err
			}

			before, err := strconv.Atoi(pages[1])
			if err != nil {
				return "", err
			}

			if rules[what] != nil {
				rules[what] = append(rules[what], before)
			} else {
				rules[what] = []int{before}

			}

		} else if u {
			pages := strings.Split(line, ",")

			var newRoot page
			newRoot.setParent(&newRoot)

			for _, p := range pages {
				what, err := strconv.Atoi(p)
				if err != nil {
					continue
				}

				if newRoot.number == 0 {
					newRoot.number = what
				} else {
					newRoot.setChild(&page{number: what})
				}

			}

			updates = append(updates, update{newRoot})

		}
	}

	sum := 0
	for _, up := range updates {
		var seen []int
		invalid := false

		for _, step := range up.root.getSteps() {
			current := step.number
			for _, rule := range rules[current] {
				if invalid {
					break
				}
				for _, s := range seen {
					if s == rule {
						invalid = true
						break
					}
				}
			}
			if invalid {
				break
			}
			seen = append(seen, current)
		}
		if !invalid {
			valid = append(valid, up)
			sum += up.root.getSteps()[(len(up.root.getSteps())-1)/2].number
		}
	}

	return strconv.Itoa(sum), nil
}

func (s Solution) Star2(input string) (string, error) {
	lines := stream.ReadLines(input)

	rules := make(map[int][]int)
	var updates []update
	var valid []update

	for _, line := range lines {
		r := false
		u := false
		for _, rune := range line {
			if string(rune) == "|" {
				r = true
				break
			} else if string(rune) == "," {
				u = true
				break
			}
		}

		if r {
			pages := strings.Split(line, "|")

			what, err := strconv.Atoi(pages[0])
			if err != nil {
				return "", err
			}

			before, err := strconv.Atoi(pages[1])
			if err != nil {
				return "", err
			}

			if rules[what] != nil {
				rules[what] = append(rules[what], before)
			} else {
				rules[what] = []int{before}

			}

		} else if u {
			pages := strings.Split(line, ",")

			var newRoot page
			newRoot.setParent(&newRoot)

			for _, p := range pages {
				what, err := strconv.Atoi(p)
				if err != nil {
					continue
				}

				if newRoot.number == 0 {
					newRoot.number = what
				} else {
					newRoot.setChild(&page{number: what})
				}

			}

			updates = append(updates, update{newRoot})

		}
	}

	sum := 0
	for _, up := range updates {
		var seen []int
		invalid := false

		for _, step := range up.root.getSteps() {
			current := step.number
			for _, rule := range rules[current] {
				if invalid {
					break
				}
				for _, s := range seen {
					if s == rule {
						invalid = true
						break
					}
				}
			}
			if invalid {
				break
			}
			seen = append(seen, current)
		}
		if invalid {
			valid = append(valid, up)
		}
	}

	for _, up := range valid {

	}

	return strconv.Itoa(sum), nil
}
