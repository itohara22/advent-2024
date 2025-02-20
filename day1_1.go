package main

import (
	"fmt"
	"math"
	"slices"
)

func day1() {
	left, right := twoArraysFromInput("day1.txt")
	slices.Sort(left)
	slices.Sort(right)

	diffs := []int{}
	for i := 0; i < len(left); i++ {
		l_val := float64(left[i])
		r_val := float64(right[i])
		diffs = append(diffs, int(math.Abs(l_val-r_val)))
	}

	res := 0
	for _, val := range diffs {
		res += val
	}
	fmt.Println(res)
}
