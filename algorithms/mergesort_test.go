package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{7, 99, 10, 2, 4, 100, 1, 1, 99},
			output: []int{1, 1, 2, 4, 7, 10, 99, 99, 100},
		},
		{
			input:  []int{5},
			output: []int{5},
		},
		{
			input:  []int{},
			output: []int{},
		},
	}

	for _, c := range cases {
		result := MergeSort(c.input)
		assert.Equal(t, c.output, result, "MergeSort(%v) == %v, want %v", c.input, result, c.output)
	}
}
