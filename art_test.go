package ghart

import (
	"testing"
	"time"
)

func TestPointToDate(t *testing.T) {
	testCases := []struct {
		p    Point
		date time.Time
	}{
		// {Point{0, 5}, time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local)},
		// {Point{0, 6}, time.Date(2021, time.January, 2, 0, 0, 0, 0, time.Local)},
		// {Point{1, 0}, time.Date(2021, time.January, 3, 0, 0, 0, 0, time.Local)},
		// {Point{1, 1}, time.Date(2021, time.January, 4, 0, 0, 0, 0, time.Local)},

		// {Point{5, 1}, time.Date(2021, time.February, 1, 0, 0, 0, 0, time.Local)},
		// {Point{52, 5}, time.Date(2021, time.December, 31, 0, 0, 0, 0, time.Local)},
	}

	for _, tc := range testCases {
		if res := pointToDate(tc.p); res != tc.date {
			t.Errorf("For point %v, date: %v, want: %v", tc.p, res, tc.date)
		}
	}
}
