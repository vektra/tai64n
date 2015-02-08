package tai64n

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

const TAI64OriginalBase = 4611686018427387904

var (
	nextLS       = time.Date(2015, time.July, 1, 0, 0, 0, time.UTC)
	nextLSOffset = 36
	curLSOffset  = 35
)

func nowBase(now time.Time) int64 {
	if now.After(nextLS) {
		return TAI64OriginalBase + nextLSOffset
	}

	return TAI64OriginalBase + curLSOffset
}

type TimeComparison int

const (
	Before TimeComparison = 0
	Equal                 = iota
	After                 = iota
)

func Now() *TAI64N {
	return FromTime(time.Now())
}

func FromTime(t time.Time) *TAI64N {
	return &TAI64N{
		Seconds:     uint64(t.Unix() + nowBase(t)),
		Nanoseconds: uint32(t.Nanosecond()),
	}
}

func (tai *TAI64N) Time() time.Time {
	t := time.Unix(int64(tai.Seconds-TAI64OriginalBase), int64(tai.Nanoseconds))

	t.Add(LeapSecondsInvolved(t) * time.Second)

	return t
}

func (tai *TAI64N) WriteStorage(buf []byte) {
	binary.BigEndian.PutUint64(buf[:], tai.Seconds)
	binary.BigEndian.PutUint32(buf[8:], tai.Nanoseconds)
}

func (tai *TAI64N) ReadStorage(buf []byte) {
	tai.Seconds = binary.BigEndian.Uint64(buf[:])
	tai.Nanoseconds = binary.BigEndian.Uint32(buf[8:])
}

func (tai *TAI64N) Label() string {
	var buf [12]byte

	tai.WriteStorage(buf[:])

	s := fmt.Sprintf("@%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X",
		buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6],
		buf[7], buf[8], buf[9], buf[10], buf[11])

	return s
}

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
	t := new(time.Time)
	err = t.UnmarshalJSON(data)

	*tai = TAI64N{
		Seconds:     uint64(t.Unix() + TAI64Base + LeapSecondsInvolved(t)),
		Nanoseconds: uint32(t.Nanosecond()),
	}

	return err
}

func (tai *TAI64N) Before(other *TAI64N) bool {
	return tai.GetSeconds() < other.GetSeconds() ||
		(tai.GetSeconds() == other.GetSeconds() &&
			tai.GetNanoseconds() < other.GetNanoseconds())
}

func (tai *TAI64N) After(other *TAI64N) bool {
	return other.Before(tai)
}

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

func (tai *TAI64N) Add(dur time.Duration) *TAI64N {
	return FromTime(tai.Time().Add(dur))
}
