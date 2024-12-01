package puzzles

import (
	y2023 "github.com/IvanKuzyshyn/aoc-go/puzzles/2023"
	y2024 "github.com/IvanKuzyshyn/aoc-go/puzzles/2024"
	"github.com/IvanKuzyshyn/aoc-go/solver"
)

type puzzle struct {
	Day      int16
	Year     int16
	Solution solver.Solver
}

var (
	allPuzzles = []puzzle{
		{
			Day:      1,
			Year:     2023,
			Solution: y2023.Day1Solution{},
		},
		// Year 2024
		{
			Day:      1,
			Year:     2024,
			Solution: y2024.Day1Solution{},
		},
	}
)

func NewPuzzlesRegistry() *solver.Registry {
	registry := solver.NewRegistry()
	for _, pzl := range allPuzzles {
		registry.Register(pzl.Year, pzl.Day, pzl.Solution)
	}
	return registry
}
