package timestamp_test

import (
	"testing"

	"gameserver/timestamp"

	"github.com/stretchr/testify/assert"
)

func TestTimestampYear4IsLeapYear(t *testing.T) {
	var ts timestamp.Timestamp
	for {
		year := ts.Year()
		if year == 4 {
			assert.True(t, ts.IsLeapYear())
			break
		} else if year > 4 {
			break
		}

		ts = ts.Next()
	}
}

func TestTimestampYear100IsNotLeapYear(t *testing.T) {
	var ts timestamp.Timestamp
	for {
		year := ts.Year()
		if year == 100 {
			assert.False(t, ts.IsLeapYear())
			break
		} else if year > 100 {
			break
		}

		ts = ts.Next()
	}
}

func TestTimestampYear400IsLeapYear(t *testing.T) {
	var ts timestamp.Timestamp
	for {
		year := ts.Year()
		if year == 400 {
			assert.True(t, ts.IsLeapYear())
		} else if year > 400 {
			break
		}

		ts = ts.Next()
	}
}

func TestTimestampYear1000IsNotLeapYear(t *testing.T) {
	var ts timestamp.Timestamp
	for {
		year := ts.Year()
		if year == 1000 {
			assert.False(t, ts.IsLeapYear())
		} else if year > 1000 {
			break
		}

		ts = ts.Next()
	}
}

func TestLeapYearFebruaryHas29Days(t *testing.T) {
	var ts timestamp.Timestamp
	for {
		year, month, day := ts.Date()
		if year == 4 && month == timestamp.February {
			assert.LessOrEqual(t, day, 29)
		} else if year > 4 {
			break
		}

		ts = ts.Next()
	}
}

func TestLeapYearHasAtMost366Days(t *testing.T) {
	var ts timestamp.Timestamp
	for {
		year := ts.Year()
		yday := ts.YearDay()
		if year == 4 {
			assert.LessOrEqual(t, yday, 366)
		} else if year > 4 {
			break
		}

		ts = ts.Next()
	}
}

func TestNormalYearHasAtMost365Days(t *testing.T) {
	var ts timestamp.Timestamp
	for {
		year := ts.Year()
		yday := ts.YearDay()
		if year == 1 {
			assert.LessOrEqual(t, yday, 365)
		} else if year > 1 {
			break
		}

		ts = ts.Next()
	}
}
