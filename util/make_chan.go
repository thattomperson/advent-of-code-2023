package util

import "log"

func MakeChan[V any](values []V) <-chan V {
	out := make(chan V, 2000)

	go func() {
		for i, v := range values {
			if i%10000 == 0 {
				log.Printf("Done %d %d%%", i, i/len(values))
			}

			out <- v
		}

		close(out)
	}()

	return out
}
