package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterMap(t *testing.T) {
	type args struct {
		values []uint
		filter func(uint) bool
		mapper func(uint) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "OK",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				filter: func(u uint) bool { return u%2 == 0 },
				mapper: func(u uint) string { return fmt.Sprintf("%d", u) },
			},
			want: []string{"0", "2", "4", "6", "8"},
		},
		{
			name: "no values",
			args: args{
				filter: func(u uint) bool { return u%2 == 0 },
				mapper: func(u uint) string { return fmt.Sprintf("%d", u) },
			},
			want: nil,
		},
		{
			name: "all filtered",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				filter: func(u uint) bool { return u > 10 },
				mapper: func(u uint) string { return fmt.Sprintf("%d", u) },
			},
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := FilterMap(testCase.args.values, testCase.args.filter, testCase.args.mapper)
			assert.Equal(t, testCase.want, result)
		})
	}
}
