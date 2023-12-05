package main

import (
	"strconv"
	"strings"

	"ttp.sh/advent-of-code/2023/util"
)

const numbers = "1234567890"

func Run(lines <-chan string) int {
	sum := 0
	for line := range lines {
		first := line[strings.IndexAny(line, numbers)]
		last := line[strings.LastIndexAny(line, numbers)]

		valueAsString := string(first) + string(last)
		value, _ := strconv.Atoi(valueAsString)
		sum += value
	}

	return sum
}

func main() {
	println(Run(util.Read("./input.txt")))
}
