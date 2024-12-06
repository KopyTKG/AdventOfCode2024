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

func (u *update) valid(rules map[int][]int) bool {
	var seen []int

	for _, step := range (*u).root.getSteps() {
		current := step.number
		for _, rule := range rules[current] {
			for _, s := range seen {
				if s == rule {
					return false
				}
			}
		}
		seen = append(seen, current)
	}
	return true

}

func (s Solution) Star1(input string) (string, error) {
	lines := stream.ReadLines(input)

	rules := make(map[int][]int)
	var updates []update

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
		if up.valid(rules) {
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
		if !up.valid(rules) {
			valid = append(valid, up)
		}
	}

	allValid := false
	for !allValid {
		allValid = true
		for pos, up := range valid {
			if up.valid(rules) {
				continue
			}
			allValid = false
			var row []int
			for _, item := range up.root.getSteps() {
				row = append(row, item.number)
			}

			var sorted []int
			for _, x := range row {
				if len(sorted) == 0 {
					sorted = append(sorted, x)
				} else {
					inserted := false
					for _, rule := range rules[x] {
						for i, k := range sorted {
							if k == rule {
								sorted = append(sorted[:i], append([]int{x}, sorted[i:]...)...)
								inserted = true
								break
							}
						}
						if inserted {
							break
						}
					}
					if !inserted {
						sorted = append(sorted, x)
					}
				}
			}

			var test update
			test.root = page{number: sorted[0]}
			test.root.setParent(&test.root)

			for i, x := range sorted {
				if i == 0 {
					continue
				}

				test.root.setChild(&page{number: x})
			}

			valid[pos] = test
		}
	}

	for _, up := range valid {
		var row []int
		for _, item := range up.root.getSteps() {
			row = append(row, item.number)
		}
		if len(row)%2 == 0 {
			sum += row[len(row)/2]
		} else {
			sum += row[(len(row)-1)/2]
		}

	}
	return strconv.Itoa(sum), nil
}
