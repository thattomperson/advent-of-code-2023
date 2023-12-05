package main

import (
	"math"
	"strconv"
	"strings"

	"ttp.sh/advent-of-code/2023/util"
)

var numbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func FindFirst(str string, matches []string) int {
	index := math.MaxInt
	res := -1
	for i, v := range matches {
		new_index := strings.Index(str, v)
		if new_index < index && new_index > -1 {
			res = i
			index = new_index
		}
		new_index = strings.Index(str, strconv.Itoa(i))
		if new_index < index && new_index > -1 {
			res = i
			index = new_index
		}
	}
	return res
}

func FindLast(str string, matches []string) int {
	index := math.MinInt
	res := -1
	for i, v := range matches {
		new_index := strings.LastIndex(str, v)
		if new_index > index && new_index > -1 {
			res = i
			index = new_index
		}
		new_index = strings.LastIndex(str, strconv.Itoa(i))
		if new_index > index && new_index > -1 {
			res = i
			index = new_index
		}
	}
	return res
}

func Run(lines <-chan string) int {
	sum := 0
	for line := range lines {
		first := FindFirst(line, numbers)
		last := FindLast(line, numbers)
		sum += (first * 10) + last
	}

	return sum
}

func main() {
	println(Run(util.Read("./input.txt")))
}
