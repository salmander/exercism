package clock

import "fmt"

const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	Hour   int
	Minute int
}

func New(hour, minute int) Clock {
	// Add the minutes to the hour
	hour += (minute / 60)

	// Wrap around
	hour = hour % 24
	minute = minute % 60

	// Check and fix if the hour is negative
	fixNegativeHours(&hour)

	// Check and fix if the minute is negative
	fixNegativeMins(&minute, &hour)

	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	// Add the total minutes
	minutes = c.Minute + minutes
	hours := (c.Hour + minutes/60)

	// Check and fix if the minute is negative
	fixNegativeMins(&minutes, &hours)

	// Check and fix if the hour is negative
	fixNegativeHours(&hours)

	c.Hour = hours % 24
	c.Minute = minutes % 60
	return c
}

// Check and fix if the hour is negative
func fixNegativeHours(hour *int) {
	if *hour < 0 {
		*hour = (24 + *hour) % 24
	}
}

// Check and fix if the minute is negative
func fixNegativeMins(minutes *int, hours *int) {
	if *minutes < 0 {
		*minutes = 60 + *minutes%60
		*hours--
	}
}
