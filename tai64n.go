package tai64n

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

// The number of seconds to offset the UNIX epoch to find the current
// TAI time.
const TAI64OriginalBase = uint64(4611686018427387904)

var (
	curLS       = time.Date(2015, time.July, 1, 0, 0, 0, 0, time.UTC)
	curLSOffset = uint64(36)
)

func nowBase(now time.Time) int64 {
	sec := now.Unix()

	// perf bias: most users set their server time to the current
	// time on earth, so we bias this to check that we're in that
	// time region before checking the complete leap second table.

	switch {
	case sec >= curLS.Unix():
		return int64(TAI64OriginalBase + curLSOffset)
	default:
		return int64(TAI64OriginalBase + LeapSecondsInvolved(now))
	}
}

// Indicates via Before, After, or Equal how to moments compare to eachother.
type TimeComparison int

const (
	Before TimeComparison = -1
	Equal  TimeComparison = 0
	After  TimeComparison = 1
)

// Return the current moment
func Now() *TAI64N {
	t := time.Now()

	return &TAI64N{
		Seconds:     uint64(t.Unix() + nowBase(t)),
		Nanoseconds: uint32(t.Nanosecond()),
	}
}

// Convert from a time.Time
func FromTime(t time.Time) *TAI64N {
	return &TAI64N{
		Seconds:     uint64(t.Unix() + int64(TAI64OriginalBase+LeapSecondsInvolved(t))),
		Nanoseconds: uint32(t.Nanosecond()),
	}
}

// Convert back to a time.Time
func (tai *TAI64N) Time() time.Time {
	t := time.Unix(int64(tai.Seconds-TAI64OriginalBase), int64(tai.Nanoseconds)).UTC()

	return t.Add(-time.Duration(LeapSecondsInvolved(t)) * time.Second)
}

// Returns `true` if TAI64N represents the zero time instant
func (t *TAI64N) IsZero() bool {
	return t.Time().IsZero()
}

// Return the value in it's canonical binary format
func (tai *TAI64N) WriteStorage(buf []byte) {
	binary.BigEndian.PutUint64(buf[:], tai.Seconds)
	binary.BigEndian.PutUint32(buf[8:], tai.Nanoseconds)
}

// Update the value from it's canonical binary format
func (tai *TAI64N) ReadStorage(buf []byte) {
	tai.Seconds = binary.BigEndian.Uint64(buf[:])
	tai.Nanoseconds = binary.BigEndian.Uint32(buf[8:])
}

// Render the moment in the canonical ascii format
func (tai *TAI64N) Label() string {
	var buf [12]byte

	tai.WriteStorage(buf[:])

	s := fmt.Sprintf("@%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X",
		buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6],
		buf[7], buf[8], buf[9], buf[10], buf[11])

	return s
}

// Parse the canonical ascii format
func ParseTAI64NLabel(label string) *TAI64N {
	if label[0] != '@' {
		return nil
	}

	buf, err := hex.DecodeString(label[1:])

	if len(buf) != 12 || err != nil {
		return nil
	}

	ts := &TAI64N{}

	ts.ReadStorage(buf[:])

	return ts
}

func (tai TAI64N) MarshalJSON() ([]byte, error) {
	return tai.Time().MarshalJSON()
}

func (tai *TAI64N) UnmarshalJSON(data []byte) (err error) {
	var t time.Time
	err = t.UnmarshalJSON(data)

	*tai = *FromTime(t)

	return err
}

// Indicated if the called moment is before the argument
func (tai *TAI64N) Before(other *TAI64N) bool {
	return tai.Compare(other) == Before
}

// Indicated if the called moment is after the argument
func (tai *TAI64N) After(other *TAI64N) bool {
	return tai.Compare(other) == After
}

// Indicate how the 2 moments compare to eachother
func (tai *TAI64N) Compare(other *TAI64N) TimeComparison {
	if tai.Seconds < other.Seconds {
		return Before
	}

	if tai.Seconds > other.Seconds {
		return After
	}

	if tai.Nanoseconds < other.Nanoseconds {
		return Before
	}

	if tai.Nanoseconds > other.Nanoseconds {
		return After
	}

	return Equal
}

// Generate a new moment by adding a duration
func (tai *TAI64N) Add(dur time.Duration) *TAI64N {
	var (
		secs  = int64(tai.Seconds) + int64(dur/time.Second)
		nsecs = int32(tai.Nanoseconds) + int32(dur%time.Second)
	)

	if nsecs >= 1e9 {
		secs++
		nsecs -= 1e9
	} else if nsecs < 0 {
		secs--
		nsecs += 1e9
	}

	return &TAI64N{
		Seconds:     uint64(secs),
		Nanoseconds: uint32(nsecs),
	}
}

// Return a duration as the difference between the 2 times
func (tai *TAI64N) Sub(other *TAI64N) time.Duration {
	return tai.Time().Sub(other.Time())
}
