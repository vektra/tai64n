package tai64n

import "time"

// Represents the first moment after a leap second occurs.
type LeapSecond struct {
	Threshold time.Time
	Offset    int
}

var AllLeapSeconds = []*LeapSecond{
	&LeapSecond{time.Date(1972, time.January, 1, 0, 0, 0, 0, time.UTC), 10},
	&LeapSecond{time.Date(1972, time.July, 1, 0, 0, 0, 0, time.UTC), 11},
	&LeapSecond{time.Date(1973, time.January, 1, 0, 0, 0, 0, time.UTC), 12},
	&LeapSecond{time.Date(1974, time.January, 1, 0, 0, 0, 0, time.UTC), 13},
	&LeapSecond{time.Date(1975, time.January, 1, 0, 0, 0, 0, time.UTC), 14},
	&LeapSecond{time.Date(1976, time.January, 1, 0, 0, 0, 0, time.UTC), 15},
	&LeapSecond{time.Date(1977, time.January, 1, 0, 0, 0, 0, time.UTC), 16},
	&LeapSecond{time.Date(1978, time.January, 1, 0, 0, 0, 0, time.UTC), 17},
	&LeapSecond{time.Date(1979, time.January, 1, 0, 0, 0, 0, time.UTC), 18},
	&LeapSecond{time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC), 19},
	&LeapSecond{time.Date(1981, time.July, 1, 0, 0, 0, 0, time.UTC), 20},
	&LeapSecond{time.Date(1982, time.July, 1, 0, 0, 0, 0, time.UTC), 21},
	&LeapSecond{time.Date(1983, time.July, 1, 0, 0, 0, 0, time.UTC), 22},
	&LeapSecond{time.Date(1985, time.July, 1, 0, 0, 0, 0, time.UTC), 23},
	&LeapSecond{time.Date(1988, time.January, 1, 0, 0, 0, 0, time.UTC), 24},
	&LeapSecond{time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC), 25},
	&LeapSecond{time.Date(1991, time.January, 1, 0, 0, 0, 0, time.UTC), 26},
	&LeapSecond{time.Date(1992, time.July, 1, 0, 0, 0, 0, time.UTC), 27},
	&LeapSecond{time.Date(1993, time.July, 1, 0, 0, 0, 0, time.UTC), 28},
	&LeapSecond{time.Date(1994, time.July, 1, 0, 0, 0, 0, time.UTC), 29},
	&LeapSecond{time.Date(1996, time.January, 1, 0, 0, 0, 0, time.UTC), 30},
	&LeapSecond{time.Date(1997, time.July, 1, 0, 0, 0, 0, time.UTC), 31},
	&LeapSecond{time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC), 32},
	&LeapSecond{time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC), 33},
	&LeapSecond{time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC), 34},
	&LeapSecond{time.Date(2012, time.July, 1, 0, 0, 0, 0, time.UTC), 35},
	&LeapSecond{time.Date(2015, time.July, 1, 0, 0, 0, 0, time.UTC), 36},
	&LeapSecond{time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC), 37},
}

type LeapMoment struct {
	LeapSecond *LeapSecond
	Moment     *TAI64N
}

var AllLeapMoments []*LeapMoment

func init() {
	for _, ls := range AllLeapSeconds {
		moment := FromTime(ls.Threshold)
		moment.Seconds--

		AllLeapMoments = append(AllLeapMoments, &LeapMoment{ls, moment})
	}
}

// Return the number of leap seconds that occur previous to the given
// time.
func LeapSecondsInvolved(t time.Time) uint64 {
	// performance bias: typically times will be in the recent history,
	// because, well, computers. So check from most recent leap second
	// backwards.

	for i := len(AllLeapSeconds) - 1; i >= 0; i-- {
		ls := AllLeapSeconds[i]
		if t.Unix() >= ls.Threshold.Unix() {
			return uint64(ls.Offset)
		}
	}

	return 0
}

func nearestLeapMoment(t *TAI64N) *LeapMoment {
	for i := len(AllLeapMoments) - 1; i >= 0; i-- {
		lm := AllLeapMoments[i]

		if t.Equal(lm.Moment) || t.After(lm.Moment) {
			return lm
		}
	}

	return nil
}
