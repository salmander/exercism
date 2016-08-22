package hamming

import "errors"

const testVersion = 4

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return -1, errors.New("Different size")
	}

	for i, letterA := range a {
		letterB := rune(b[i])
		if letterA != letterB {
			count++
		}
	}
	return count, nil
}
