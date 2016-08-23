package raindrops

import "strconv"

const testVersion = 2

var numbersToTest = []int{3, 5, 7}

func Convert(number int) string {
	return getRainDropsFor(number)
}

func getRainDropsFor(number int) string {
	var words string
	for _, i := range numbersToTest {
		if number%i == 0 {
			words += numberToString(i)
		}
	}

	// Check if there is nothing in the "words" string,
	// then return the original number
	if words == "" {
		words += strconv.Itoa(number)
	}
	return words
}

func numberToString(number int) string {
	switch number {
	case 3:
		return "Pling"
	case 5:
		return "Plang"
	case 7:
		return "Plong"
	default:
		return ""
	}
}

// The test program has a benchmark too.  How fast does your Convert convert?
// go test -bench .
// BenchmarkConvert-4       1000000              1061 ns/op
