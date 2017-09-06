package clock

import (
	"fmt"
)

const testVersion = 4

// You can find more details and hints in the test file.

type Clock struct {
	hour   int
	minute int
}

// 1. for positive hrs and mins.
// a. start with mins then hours
// - if more than 59 mins:
//   - hrs = mins/60 and mins = mins%60 eg for 1723mins, then hrs=1723/60 and mins=1723%60
// - if more than 24 hours:
//   - find how many hours are above 24 and use that as the hour.
// 	if hour =27, then hour = 27-24=03 == 27%24
// 	if hour=100, then hour  = 100%24 = 4

// 2. for negative hrs and/or mins
// b. start with mins then helpers
// - for negative mins
//   hrs = (hrs + -mins/60)%24 and mins = -mins%60 eg for -160mins(hrs= (hrs + -160/60)%24 and mins= -40%60)
// - for negative hrs
//   hrs = -hrs%24 eg (-91%24=05)
func positive_modulo(m, n int) int {
	return ((m % n) + n) % n
}

func New(hour, minute int) Clock {
	fmt.Println("1. hour, min", hour, minute)
	if minute < 0 {
		hour = positive_modulo((hour + minute/60), 24)
		minute = positive_modulo(minute, 60)
	}

	if hour < 0 {
		fmt.Println("2. hour, min", hour, minute)
		hour = positive_modulo(hour, 24)
		fmt.Println("3. hour, min", hour, minute)
	}
	fmt.Println("4. hour, min", hour, minute)
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

func (c Clock) Add(minutes int) Clock {
	clockinmins := c.clockinmins()
	clockinmins = clockinmins + minutes

	numberofhrs := clockinmins / 60
	numberofmins := clockinmins % 60
	c.hour = numberofhrs
	c.minute = numberofmins

	return c
}

/* clock := New(10, 30)
// subtract an hour and a half from it
clock = clock.Add(-90)
fmt.Println(clock.String())
Output: 09:00
*/
// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
