package y2023

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"unicode"

	"github.com/IvanKuzyshyn/aoc-go/solver"
)

type Day1Solver struct{}

type calibrator struct {
	rows  chan string
	wg    sync.WaitGroup
	total int64
}

func (c *calibrator) addRow(row string) {
	if len(row) == 0 {
		fmt.Println("Skipped row")
		// Ignore empty rows
		return
	}
	c.rows <- row
}

func (c *calibrator) startCalculation() {
	for r := range c.rows {
		c.wg.Add(1)
		go c.calcRowCalibrationVal(r)
	}
}

func (c *calibrator) calcRowCalibrationVal(row string) {
	defer c.wg.Done()
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
	total, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		// If one row fails the whole result fails
		panic(err)
	} else {
		atomic.AddInt64(&c.total, int64(total))
	}
}

func (s Day1Solver) Solve(opts solver.Opts) (solver.Result, error) {
	c := calibrator{
		rows:  make(chan string),
		total: 0,
	}
	rows := strings.Split(opts.Input, "\n")

	go c.startCalculation()

	for _, row := range rows {
		c.addRow(row)
	}
	close(c.rows)
	c.wg.Wait()

	return solver.Result{
		Output: strconv.Itoa(int(c.total)),
	}, nil
}
