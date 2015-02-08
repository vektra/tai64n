package tai64n

import "time"

type LeapSecond struct {
	Threshold time.Time
	Offset    int
}

var AllLeapSeconds = []*LeapSecond{
	&LeapSecond{time.Date(1972, 1, 1, 0, 0, 0, time.UTC), 10},
	&LeapSecond{time.Date(1972, 7, 1, 0, 0, 0, time.UTC), 11},
	&LeapSecond{time.Date(1973, 1, 1, 0, 0, 0, time.UTC), 12},
	&LeapSecond{time.Date(1974, 1, 1, 0, 0, 0, time.UTC), 13},
	&LeapSecond{time.Date(1975, 1, 1, 0, 0, 0, time.UTC), 14},
	&LeapSecond{time.Date(1976, 1, 1, 0, 0, 0, time.UTC), 15},
	&LeapSecond{time.Date(1977, 1, 1, 0, 0, 0, time.UTC), 16},
	&LeapSecond{time.Date(1978, 1, 1, 0, 0, 0, time.UTC), 17},
	&LeapSecond{time.Date(1979, 1, 1, 0, 0, 0, time.UTC), 18},
	&LeapSecond{time.Date(1980, 1, 1, 0, 0, 0, time.UTC), 19},
	&LeapSecond{time.Date(1981, 7, 1, 0, 0, 0, time.UTC), 20},
	&LeapSecond{time.Date(1982, 7, 1, 0, 0, 0, time.UTC), 21},
	&LeapSecond{time.Date(1983, 7, 1, 0, 0, 0, time.UTC), 22},
	&LeapSecond{time.Date(1985, 7, 1, 0, 0, 0, time.UTC), 23},
	&LeapSecond{time.Date(1988, 1, 1, 0, 0, 0, time.UTC), 24},
	&LeapSecond{time.Date(1990, 1, 1, 0, 0, 0, time.UTC), 25},
	&LeapSecond{time.Date(1991, 1, 1, 0, 0, 0, time.UTC), 26},
	&LeapSecond{time.Date(1992, 7, 1, 0, 0, 0, time.UTC), 27},
	&LeapSecond{time.Date(1993, 7, 1, 0, 0, 0, time.UTC), 28},
	&LeapSecond{time.Date(1994, 7, 1, 0, 0, 0, time.UTC), 29},
	&LeapSecond{time.Date(1996, 1, 1, 0, 0, 0, time.UTC), 30},
	&LeapSecond{time.Date(1997, 7, 1, 0, 0, 0, time.UTC), 31},
	&LeapSecond{time.Date(1999, 1, 1, 0, 0, 0, time.UTC), 32},
	&LeapSecond{time.Date(2006, 1, 1, 0, 0, 0, time.UTC), 33},
	&LeapSecond{time.Date(2009, 1, 1, 0, 0, 0, time.UTC), 34},
	&LeapSecond{time.Date(2012, 7, 1, 0, 0, 0, time.UTC), 35},
	&LeapSecond{time.Date(2015, 7, 1, 0, 0, 0, time.UTC), 36},
}

func LeapSecondsInvolved(t time.Time) int {
	// performance bias: typically times will be in the recent history,
	// because, well, computers. So check from most recent leap second
	// backwards.

	for i := len(AllLeapSeconds) - 1; i >= 0; i-- {
		ls := AllLeapSeconds[i]
		if t.Equal(ls.Threshold) || t.After(ls.Threshold) {
			return ls.Offset
		}
	}

	return 0
}
