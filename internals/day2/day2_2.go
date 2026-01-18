package day2

import (
	"fmt"
	"kaddu/utils"
	"math"
)

func Day2_2() {
	reports := utils.ReadFileForDay2("S:/advent-2024/internals/day2/day2.txt")
	for _,report := range reports{
		fmt.Println(report)
	}
}

func isReportValid(report []int) bool {
	error_tolerate := 1

	for i:=1, i< len(report); i++ {
	diff := math.Abs(int(report[i]) - int(report[i-1]))
	}
}

