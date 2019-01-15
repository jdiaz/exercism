// Package clock contains functions pertaining to the implementation of a time-only clock.
package clock

import (
	"fmt"
)

//Clock struct represents a time-only clock (hour, minutes)
type Clock struct {
	hour, minutes int
}

// New initializes a clock with default hour and minutes values
func New(hr, min int) Clock {
	h, m := calculateTime(hr, min)
	return Clock{hour: h, minutes: m}
}

// String returns a string representation of the clock
func (c Clock) String() string {
	hour := "%d"
	minutes := "%d"

	if isSingleDigit(c.hour) {
		hour = "0%d"
	}

	if isSingleDigit(c.minutes) {
		minutes = "0%d"
	}

	return fmt.Sprintf(hour+":"+minutes, c.hour, c.minutes)
}

// Add increments the time by the provides minutes amount
func (c Clock) Add(minutes int) Clock {
	h, m := calculateTime(c.hour, c.minutes+minutes)
	c.hour = h
	c.minutes = m
	return c
}

// Subtract decreases the time by the provides minutes amount
func (c Clock) Subtract(minutes int) Clock {
	h, m := calculateTime(c.hour, c.minutes-minutes)
	c.hour = h
	c.minutes = m
	return c
}

func calculateTime(hour, minutes int) (int, int) {
	if hour < 0 {
		hour = -24 - hour
		hour = hour * -1
		fmt.Println(hour)
	}
	if minutes < 0 {
		minutes = -60 - minutes
		minutes = minutes * -1
		fmt.Println(minutes)
	}

	increment := minutes / 60
	return (hour + increment) % 24, minutes % 60
}

func isSingleDigit(n int) bool {
	return n/10 == 0
}
