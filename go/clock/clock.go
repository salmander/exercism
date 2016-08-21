package clock

import "fmt"

const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	Hour   int
	Minute int
} // Complete the type definition.  Pick a suitable data type.

func New(hour, minute int) Clock {
	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%v:%v", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	c.Minute += minutes
	return c
}
