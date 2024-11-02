package puzzles

import (
	"fmt"

	y2023 "github.com/IvanKuzyshyn/aoc-go/puzzles/2023"
)

type Resolver interface {
	Resolve(input string) (string, error)
}

type Solution struct {
	Day  int
	Year int
}

func (p *Solution) Solve(input string) (string, error) {
	// Temporary logic
	if p.Day != 1 || p.Year != 2023 {
		return "", fmt.Errorf("day should be %d in %d year", p.Day, p.Year)
	}

	return y2023.SolveD1(input)
}
