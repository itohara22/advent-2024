package day2

import (
	"fmt"
	"kaddu/utils"
	"slices"
)

func Day2_2() {
	reports := utils.ReadFileForDay2("S:/advent-2024/internals/day2/day2.txt")
	res := 0
	for _,report := range reports{
		if isReportValid(report) {
			res++
		}
	}
	fmt.Println(res)
}

func isReportValid(report []int) bool {
	if is_report_valid(report){
		return true
	}

	for i:=0; i< len(report); i++ {
		tmp := append( []int{}, report...)
		tmp = slices.Delete(tmp,i,i+1) // need i+1 to remove and check fo rlast element
		if is_report_valid(tmp) {
			return true
		}
	}
	return false
}

