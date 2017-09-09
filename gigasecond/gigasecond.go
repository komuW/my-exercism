package gigasecond

import "time"

const testVersion = 4

// AddGigasecond adds a gigasec to the given time
func AddGigasecond(inputTime time.Time) time.Time {
	age := inputTime.Add(time.Second * 1000000000)
	return age
}
