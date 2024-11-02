package puzzles

import (
	"strconv"
	"strings"
	"unicode"
)

func SolveD1(input string) (string, error) {
	var totalCalibrationVal int

	for _, row := range strings.Split(input, "\n") {
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
			return "", err
		}
		totalCalibrationVal += totalInRow
	}

	return strconv.Itoa(totalCalibrationVal), nil
}
