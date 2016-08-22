// Package clause.
package gigasecond

import (
	"math"
	"time"
)

// Constant declaration.
const testVersion = 4

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	seconds := math.Pow(10, 9)
	duration := time.Duration(seconds) * time.Second

	return t.Add(duration)
}
