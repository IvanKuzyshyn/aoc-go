package y2024

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/IvanKuzyshyn/aoc-go/algorithms"
	"github.com/IvanKuzyshyn/aoc-go/solver"
)

type Day1Solution struct{}

var solutions = map[int8]func(opts solver.Opts) (solver.Result, error){
	1: solveFirstPart,
	2: solveSecondPart,
}

func (s Day1Solution) Solve(opts solver.Opts) (solver.Result, error) {
	solution, ok := solutions[opts.Part]
	if !ok {
		return solver.Result{}, fmt.Errorf("solution %d part is not supported", opts.Part)
	}

	return solution(opts)
}

func solveFirstPart(opts solver.Opts) (solver.Result, error) {
	listDist := struct {
		left  []int
		right []int
		total int
	}{
		left:  make([]int, 0),
		right: make([]int, 0),
		total: 0,
	}

	left, right, err := extractLists(opts.Input)
	if err != nil {
		return solver.Result{}, err
	}
	if len(left) != len(right) {
		return solver.Result{}, fmt.Errorf("lists have different length: %d for the left and %d for the right", len(left), len(right))
	}

	listDist.left = algorithms.MergeSort(left)
	listDist.right = algorithms.MergeSort(right)

	for i := 0; i < len(listDist.left); i++ {
		lv := listDist.left[i]
		rv := listDist.right[i]

		listDist.total += diff(lv, rv)
	}

	return solver.Result{
		Output: strconv.Itoa(listDist.total),
	}, nil
}

func solveSecondPart(opts solver.Opts) (solver.Result, error) {
	similarityScore := struct {
		left                []int
		right               []int
		rightInstancesCount map[int]int
		total               int
	}{
		left:                make([]int, 0),
		right:               make([]int, 0),
		rightInstancesCount: map[int]int{},
		total:               0,
	}

	left, right, err := extractLists(opts.Input)
	if err != nil {
		return solver.Result{}, err
	}
	if len(left) != len(right) {
		return solver.Result{}, fmt.Errorf("lists have different length: %d for the left and %d for the right", len(left), len(right))
	}
	similarityScore.left = left
	similarityScore.right = right

	for _, v := range similarityScore.right {
		similarityScore.rightInstancesCount[v]++
	}

	for _, v := range similarityScore.left {
		mult := 0
		if iv, ok := similarityScore.rightInstancesCount[v]; ok {
			mult = iv
		}

		similarityScore.total += v * mult
	}

	return solver.Result{
		Output: strconv.Itoa(similarityScore.total),
	}, nil
}

func extractLists(input string) ([]int, []int, error) {
	var left []int
	var right []int
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		fields := strings.Fields(row)
		lv, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, err
		}
		rv, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, err
		}
		left = append(left, lv)
		right = append(right, rv)
	}

	return left, right, nil
}

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
