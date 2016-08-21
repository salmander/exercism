package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// Check if the given year is a leap year
func IsLeapYear(year int) bool {
	if year < 1 {
		panic("ahhh I don't know what you mean")
	}

	a := year % 4
	b := year % 100
	c := year % 400

	// The year is divisble by 4 but not by 100
	if a == 0 && b != 0 {
		return true
	} else if a == 0 && b == 0 && c == 0 { // Except it is also divisible by 400
		return true
	}

	return false
}
