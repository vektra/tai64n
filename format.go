package tai64n

import (
	"fmt"
	"time"
)

func (t *TAI64N) Date() (year int, month time.Month, day int) {
	lm := nearestLeapMoment(t)

	if t.Equal(lm.Moment) {
		prev := lm.LeapSecond.Threshold.Add(-1 * time.Second)
		return prev.Date()
	}

	return t.Time().Date()
}

func (t *TAI64N) Clock() (hour, min, sec int) {
	lm := nearestLeapMoment(t)

	if t.Equal(lm.Moment) {
		prev := lm.LeapSecond.Threshold.Add(-1 * time.Second)
		hour, min, sec := prev.Clock()
		return hour, min, sec + 1
	}

	return t.Time().Clock()

}

func (t *TAI64N) String() string {
	year, month, day := t.Date()
	hour, minute, sec := t.Clock()

	return fmt.Sprintf("%4d-%02d-%02dT%02d:%02d:%02d.%dZ",
		year, month, day,
		hour, minute, sec,
		t.Nanoseconds)
}
