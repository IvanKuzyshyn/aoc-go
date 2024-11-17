package y2023

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/IvanKuzyshyn/aoc-go/solver"
)

type Day1Solver struct{}

func (s Day1Solver) Solve(opts solver.Opts) (solver.Result, error) {
	var totalCalibrationVal int

	for _, row := range strings.Split(opts.Input, "\n") {
		if row == "" {
			continue
		}

		var firstDigit string
		var lastDigit string

		for _, r := range row {
			if unicode.IsDigit(r) {
				if firstDigit == "" {
					firstDigit = string(r)
					continue
				}

				lastDigit = string(r)
			}
		}

		if lastDigit == "" {
			lastDigit = firstDigit
		}
		totalInRow, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			return solver.Result{}, err
		}
		totalCalibrationVal += totalInRow
	}

	return solver.Result{
		Output: strconv.Itoa(totalCalibrationVal),
	}, nil
}
