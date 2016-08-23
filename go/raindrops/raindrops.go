package raindrops

import "fmt"

const testVersion = 2

func Convert(number int) string {
	p := fmt.Println

	var output string
	var factors []int

	factors = getFactorsFor(number)

	p("Factors for number:", number, "are:", factors)

	output = numberToString(factors)

	p("Factors for numbers are:", output)

	return output
}
func getFactorsFor(number int) []int {
	var factors []int

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func numberToString(factors []int) string {
	var output string

	if len(factors) == 1 && factors[0] == 1 {
		output = "1"
	} else {
		for _, n := range factors {
			switch n {
			case 3:
				output += "Pling"
			case 5:
				output += "Plang"
			case 7:
				output += "Plong"
			}
		}
	}
	return output
}

// The test program has a benchmark too.  How fast does your Convert convert?
