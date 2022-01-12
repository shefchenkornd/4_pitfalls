package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{1, 2, 3},
			target: 3,
			want:   []int{0, 1},
		},
	}

	for _, cs := range cases {
		t.Run("twoSum", func(t *testing.T) {
			got := twoSum(cs.nums, cs.target)
			assert.Equal(t, cs.want, got)
		})
	}
}
