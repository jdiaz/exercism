// Package clock contains functions pertaining to the implementation of a time-only clock.
package clock

import (
	"fmt"
)

const hourInDay = 24
const minutesInHour = 60
const minutesInDay = minutesInHour * hourInDay

//Clock struct represents a time-only clock (hour, minutes)
type Clock struct {
	minutes int
}

// New initializes a clock with default hour and minutes values
func New(hr, mnts int) Clock {
	hr = hr % hourInDay
	mnts = mnts % minutesInDay
	if hr < 0 {
		hr = hr + hourInDay
	}
	if mnts < 0 {
		mnts = mnts + minutesInDay
	}
	minutes := hr * minutesInHour + mnts
	if minutes >= minutesInDay {
		minutes = minutes % minutesInDay
	}
	return Clock{minutes: minutes}
}

// String returns a string representation of the clock
func (c Clock) String() string {
	fmt.Println("=====")
	fmt.Println(c.minutes)
	hour := c.minutes / minutesInHour
	minutes := c.minutes % minutesInHour
	s := fmt.Sprintf("%02d:%02d", hour, minutes)
	fmt.Println(s)
	return s
}

// Add increments the time by the provides minutes amount
func (c Clock) Add(minutes int) Clock {
	newMinutes := c.minutes + minutes
	h := newMinutes / minutesInHour
	m := newMinutes % minutesInHour
	return New(h, m)
}

// Subtract decreases the time by the provides minutes amount
func (c Clock) Subtract(minutes int) Clock {
	newMinutes := c.minutes - minutes
	h := newMinutes / minutesInHour
	m := newMinutes % minutesInHour
	return New(h, m)
}
