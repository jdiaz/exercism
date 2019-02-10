// Package clock contains functions pertaining to the implementation of a time-only clock.
package clock

import "fmt"

const hourInDay = 24
const minutesInHour = 60
const minutesInDay = minutesInHour * hourInDay

//Clock struct represents a time-only clock (hour, minutes)
type Clock struct {
	minutes int
}

// New initializes a clock with default hour and minutes values
func New(hr, mnts int) Clock {
	minutes := hr*minutesInHour + mnts
	minutes = minutes % minutesInDay
	if minutes < 0 {
		minutes += minutesInDay
	}
	return Clock{minutes: minutes}
}

// String returns a string representation of the clock
func (c Clock) String() string {
	hour := c.minutes / minutesInHour
	minutes := c.minutes % minutesInHour
	return fmt.Sprintf("%02d:%02d", hour, minutes)
}

// Add increments the time by the provides minutes amount
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

// Subtract decreases the time by the provides minutes amount
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}
