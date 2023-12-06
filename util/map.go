package util

func Map[I any, O any](in []I, t func(int, I) O) []O {
	new_slice := make([]O, len(in))
	for i, v := range in {
		new_slice[i] = t(i, v)
	}
	return new_slice
}
