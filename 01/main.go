package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const numbers = "1234567890"

func main() {
	sum := 0
	for line := range read("./input.txt") {
		first := line[strings.IndexAny(line, numbers)]
		last := line[strings.LastIndexAny(line, numbers)]

		valueAsString := string(first) + string(last)
		value, _ := strconv.Atoi(valueAsString)
		sum += value
	}

	println(sum)
}

func read(path string) <-chan string {
	out := make(chan string)

	go func() {

		file, _ := os.Open(path)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- scanner.Text()
		}

		close(out)
	}()
	return out
}
