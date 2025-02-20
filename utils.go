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
