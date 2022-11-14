package timestamp

import (
	"errors"
	"fmt"
	"gameserver/proto"
)

type Timestamp int

// New creates a new timestamp from a given day and time of day, its
// output is undefined when time of day is invalid.
func New(day int, timeOfDay TimeOfDay) Timestamp {
	return Timestamp(day*4 + int(timeOfDay))
}

func FromProto(ts *proto.Timestamp) Timestamp {
	return New(int(ts.GetDay()), TimeOfDay(ts.GetTimeOfDay()))
}

func (t Timestamp) Next() Timestamp {
	return Timestamp(int(t) + 1)
}

func (t Timestamp) Date() (year int, month Month, day int) {
	year, month, day, _ = date(int(t), true)
	return
}

func (t Timestamp) IsLeapYear() bool {
	year, _, _ := t.Date()
	return isLeap(year)
}

func (t Timestamp) Year() int {
	year, _, _, _ := date(int(t), false)
	return year
}

func (t Timestamp) Month() Month {
	_, month, _, _ := date(int(t), true)
	return month
}

func (t Timestamp) Day() int {
	_, _, day, _ := date(int(t), true)
	return day
}

func (t Timestamp) YearDay() int {
	_, _, _, yday := date(int(t), false)
	return yday + 1
}

func (t Timestamp) Weekday() Weekday {
	sec := (int(t) + int(Monday)*secondsPerDay) % secondsPerWeek
	return Weekday(int(sec) / secondsPerDay)
}

func (t Timestamp) TimeOfDay() TimeOfDay {
	return TimeOfDay(t % 4)
}

func (t Timestamp) String() string {
	year, month, day := t.Date()
	return fmt.Sprintf("%04d-%02d-%02d %v", year, month, day, t.TimeOfDay())
}

var (
	ErrTimestampNegative = errors.New("timestamp must not be negative")
	ErrTimestampZero     = errors.New("timestamp must not be zero")
)

type ErrTimestampNotPositive struct {
	Timestamp Timestamp
	inner     error
}

func (e ErrTimestampNotPositive) Error() string {
	return fmt.Sprintf("timestamp must be positive: %e", e.inner)
}

func (e ErrTimestampNotPositive) Unwrap() error {
	return e.inner
}

func (t Timestamp) Validate() error {
	if t == 0 {
		return ErrTimestampNotPositive{t, ErrTimestampZero}
	}

	if t < 1 {
		return ErrTimestampNotPositive{t, ErrTimestampZero}
	}

	return nil
}
