package y2023

import (
	"testing"

	"github.com/IvanKuzyshyn/aoc-go/puzzle"
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
	solution := Day1Solution{}
	result, err := solution.Solve(puzzle.Opts{Input: input})

	assert.Nil(t, err)
	assert.Equal(t, "591", result.Output)
}
