package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func twoArraysFromInput(filename string) ([]int, []int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	left := []int{}
	right := []int{}
	s := bufio.NewScanner(f)
	for s.Scan() { // reading line by line
		a := strings.Split(s.Text(), "   ")
		n, _ := strconv.Atoi(a[0])
		m, _ := strconv.Atoi(a[1])
		left = append(left, n)
		right = append(right, m)
	}

	return left, right
}

func readFileForDay2(filename string) [][]int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	reports := [][]int{}
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " ")
		levels := []int{}
		for _, val := range a {
			num_val, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err.Error())
			}
			levels = append(levels, num_val)
		}
		reports = append(reports, levels)
	}

	return reports
}
