package solver

import (
	"fmt"
)

type Registry struct {
	solvers map[int]map[int]Solver
}

func NewRegistry() *Registry {
	return &Registry{
		solvers: make(map[int]map[int]Solver),
	}
}

func (r *Registry) Register(year, day int, solver Solver) {
	if _, exists := r.solvers[year]; !exists {
		r.solvers[year] = make(map[int]Solver)
	}
	if _, exists := r.solvers[year][day]; exists {
		// It should be fatal as we can't register multiple solvers
		panic(fmt.Sprintf("Trying to register solver for Day %d and Year %d more than once!", day, year))
	}
	r.solvers[year][day] = solver
}

func (r *Registry) GetSolver(year, day int) (Solver, error) {
	if days, exists := r.solvers[year]; exists {
		if solver, exists := days[day]; exists {
			return solver, nil
		}
	}
	return nil, fmt.Errorf("solver not found for year %d, day %d", year, day)
}