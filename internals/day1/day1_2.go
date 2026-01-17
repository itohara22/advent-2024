package day1

import (
	"fmt"
	"kaddu/utils"
)

func Day1_2() {
	left, right := utils.TwoArraysFromInput("day1.txt")

	frequency_map := map[int]int{}

	for i := 0; i < len(left); i++ {
		val_l := left[i]

		_, doExist := frequency_map[val_l]

		if !doExist {
			frequency_map[val_l] = 0
			for _, num := range right {
				if num == val_l {
					frequency_map[val_l] += 1
				}
			}
		} else {
			frequency_map[val_l] *= 2
		}

	}

	sum := 0
	for k, v := range frequency_map {
		sum += k * v
	}

	fmt.Println(sum)
}
