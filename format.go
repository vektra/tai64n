package tai64n

import (
	"fmt"
	"time"
)

// Calculate the year, month, and day of this moment. If the moment
// falls on a leap second, the displayed value will be that of the
// leap second as the 60th second of the day.
func (t *TAI64N) Date() (year int, month time.Month, day int) {
	lm := nearestLeapMoment(t)

	if t.Equal(lm.Moment) {
		prev := lm.LeapSecond.Threshold.Add(-1 * time.Second)
		return prev.Date()
	}

	return t.Time().Date()
}

// Calculate the hour, minute, and second of this moment. If the moment
// falls on a leap second, the displayed value will be that of the
// leap second as the 60th second of the day.
func (t *TAI64N) Clock() (hour, min, sec int) {
	lm := nearestLeapMoment(t)

	if t.Equal(lm.Moment) {
		prev := lm.LeapSecond.Threshold.Add(-1 * time.Second)
		hour, min, sec := prev.Clock()
		return hour, min, sec + 1
	}

	return t.Time().Clock()

}

// Render the moment as a RFC3339Nano format
func (t *TAI64N) String() string {
	year, month, day := t.Date()
	hour, minute, sec := t.Clock()

	return fmt.Sprintf("%4d-%02d-%02dT%02d:%02d:%02d.%dZ",
		year, month, day,
		hour, minute, sec,
		t.Nanoseconds)
}
