package _0759_employee_free_time

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_employeeFreeTime(t *testing.T) {
	tests := []struct {
		schedule [][]*Interval
		want     []*Interval
	}{
		{
			[][]*Interval{{{1, 2}, {5, 6}}, {{1, 3}}, {{4, 10}}},
			[]*Interval{{3, 4}},
		},
		{
			[][]*Interval{{{1, 3}, {6, 7}}, {{2, 4}}, {{2, 5}, {9, 12}}},
			[]*Interval{{5, 6}, {7, 9}},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := employeeFreeTime(tt.schedule)

			require.Equal(t, tt.want, got)
		})
	}
}
