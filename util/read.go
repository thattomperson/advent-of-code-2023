package util

import (
	"bufio"
	"os"
)

func Read(path string) <-chan string {
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
