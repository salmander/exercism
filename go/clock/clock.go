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

	//fmt.Println("hour:", hour, " minute:", minute)

	return Clock{hour % 24, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	// Check if the current time's minute and the minutes is greater than 59
	if c.Minute+minutes > 59 {
		//  Increment the hour
		c.Hour++

		c.Minute = minutes + c.Minute - 60
	} else {
		c.Minute += minutes
	}

	return c
}
