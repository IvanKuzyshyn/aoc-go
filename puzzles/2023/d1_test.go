package y2023

import (
	"testing"

	"github.com/IvanKuzyshyn/aoc-go/solver"
	"github.com/stretchr/testify/assert"
)

const (
	input = `
cmpptgjc3qhcjxcbcqgqkxhrms
9sixonefour
eighttwo2twofour9
7eightseveneightthree
tlnllks2jcfdlgsjbhpfnineone
one44fivesevenjzsfzddg
787seveneight75five
7q7nine3
hlqphphn5dtzzbqvk3three8seven
ninekkbvsfq8seven7321bpdcdh
`
)

func TestExpectedResult(t *testing.T) {
	resolver := Day1Solver{}
	result, err := resolver.Solve(solver.Opts{Input: input})

	assert.Nil(t, err)
	assert.Equal(t, "591", result.Output)
}
