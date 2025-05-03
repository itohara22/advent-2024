package main

import "fmt"

func day2_1() {
	reports := readFileForDay2("day2.txt")
	res := 0

	for _, lvls := range reports {
		if is_report_valid(lvls) {
			res += 1
		}
	}

	fmt.Println(res)
}

func is_report_valid(r []int) bool {
	is_increasing := r[1]-r[0] > 0
	len_report := len(r)

	if is_increasing {
		for i := 1; i < len_report; i++ {
			diff := r[i] - r[i-1]
			if diff > 3 || diff < 1 {
				return false
			}
		}
		return true
	} else {
		for i := 1; i < len_report; i++ {
			diff := r[i-1] - r[i]
			if diff > 3 || diff < 1 {
				return false
			}
		}
		return true
	}
}
