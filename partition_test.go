package iterago

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	type args struct {
		values    []uint
		predicate func(uint) bool
	}
	type want struct {
		validated   []uint
		invalidated []uint
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "OK",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) bool { return u%2 == 0 },
			},
			want: want{
				validated:   []uint{0, 2, 4, 6, 8},
				invalidated: []uint{1, 3, 5, 7, 9},
			},
		},
		{
			name: "OK - all validated",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) bool { return u == u },
			},
			want: want{
				validated:   []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				invalidated: nil,
			},
		},
		{
			name: "OK - all invalidated",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) bool { return u != u },
			},
			want: want{
				validated:   nil,
				invalidated: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name: "no values",
			args: args{
				values:    nil,
				predicate: func(u uint) bool { return u%2 == 0 },
			},
			want: want{
				validated:   nil,
				invalidated: nil,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			validated, invalidated := Partition(testCase.args.values, testCase.args.predicate)
			assert.Equal(t, testCase.want.validated, validated)
			assert.Equal(t, testCase.want.invalidated, invalidated)
		})
	}
}
