package util

func MakeChan[V any](values []V) <-chan V {
	out := make(chan V)

	go func() {
		for _, v := range values {
			out <- v
		}

		close(out)
	}()

	return out
}
