package clock

import (
	"fmt"
)

const testVersion = 4

// Clock holds hour and minutes .
type Clock struct {
	hour   int
	minute int
}

func positive_modulo(m, n int) int {
	return ((m % n) + n) % n
}

// New creates a new Clock
func New(hour, minute int) Clock {
	if minute < 0 {
		hour = positive_modulo((hour + minute/60), 24)
		minute = positive_modulo(minute, 60)
	}

	if hour < 0 {
		hour = positive_modulo(hour, 24)
	}
	if minute > 59 {
		hour = hour + (minute / 60)
		minute = minute % 60
	}
	if hour > 23 {
		hour = hour % 24
	}

	return Clock{hour: hour, minute: minute}
}

func (c Clock) clockinmins() int {
	return c.hour*60 + c.minute
}

func (c Clock) String() string {
	hourshaper := ""
	minuteshaper := ""
	if c.hour < 10 {
		hourshaper = "0"
	}
	if c.minute < 10 {
		minuteshaper = "0"
	}
	return fmt.Sprintf("%v%v:%v%v", hourshaper, c.hour, minuteshaper, c.minute)
}

// Add adds different clocks
func (c Clock) Add(minutes int) Clock {
	clockinmins := c.clockinmins()
	clockinmins = clockinmins + minutes

	numberofhrs := clockinmins / 60
	numberofmins := clockinmins % 60
	c.hour = numberofhrs
	c.minute = numberofmins

	return c
}
