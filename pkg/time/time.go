package time

import (
	"math"
	"time"
)

func SetClock(t time.Time, hours, minutes, seconds int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hours, minutes, seconds, 0, t.Location())
}

func GetDaysDifference(from, to time.Time) int {
	if from.Compare(to) >= 0 {
		return 0
	}

	return int(math.Ceil(to.Sub(from).Hours() / 24))
}
