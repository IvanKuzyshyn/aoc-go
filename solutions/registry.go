package solutions

import (
	"fmt"

	"github.com/IvanKuzyshyn/aoc-go/puzzle"
	y2023 "github.com/IvanKuzyshyn/aoc-go/solutions/2023"
	y2024 "github.com/IvanKuzyshyn/aoc-go/solutions/2024"
)

var (
	allSolutions = []solution{
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

type solution struct {
	Day      int16
	Year     int16
	Solution puzzle.Solver
}

type Registry struct {
	solvers map[int16]map[int16]puzzle.Solver
}

func (r *Registry) Register(year, day int16, solver puzzle.Solver) {
	if _, exists := r.solvers[year]; !exists {
		r.solvers[year] = make(map[int16]puzzle.Solver)
	}
	if _, exists := r.solvers[year][day]; exists {
		// It should be fatal as we can't register multiple solvers
		panic(fmt.Sprintf("Trying to register solver for Day %d and Year %d more than once!", day, year))
	}
	r.solvers[year][day] = solver
}

func (r *Registry) GetSolution(year, day int16) (puzzle.Solver, error) {
	if days, exists := r.solvers[year]; exists {
		if solver, exists := days[day]; exists {
			return solver, nil
		}
	}
	return nil, fmt.Errorf("solver not found for year %d, day %d", year, day)
}

func NewRegistry() *Registry {
	registry := &Registry{
		solvers: make(map[int16]map[int16]puzzle.Solver),
	}
	for _, s := range allSolutions {
		registry.Register(s.Year, s.Day, s.Solution)
	}
	return registry
}
