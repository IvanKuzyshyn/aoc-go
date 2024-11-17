package puzzles

import (
	y2023 "github.com/IvanKuzyshyn/aoc-go/puzzles/2023"
	"github.com/IvanKuzyshyn/aoc-go/solver"
)

type puzzle struct {
	Day    int
	Year   int
	Solver solver.Solver
}

var (
	allPuzzles = []puzzle{
		{
			Day:    1,
			Year:   2023,
			Solver: y2023.Day1Solver{},
		},
	}
)

func NewPuzzlesRegistry() *solver.Registry {
	registry := solver.NewRegistry()
	for _, pzl := range allPuzzles {
		registry.Register(pzl.Year, pzl.Day, pzl.Solver)
	}
	return registry
}
