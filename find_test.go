package iterago

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	type args struct {
		values    []uint
		predicate func(uint) bool
	}
	tests := []struct {
		name string
		args args
		want Option[uint]
	}{
		{
			name: "OK",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) bool { return u == 5 },
			},
			want: NewOption[uint](5),
		},
		{
			name: "not found",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) bool { return u == 666 },
			},
			want: NewNoneOption[uint](),
		},
		{
			name: "no values",
			args: args{
				values:    nil,
				predicate: func(u uint) bool { return u == 5 },
			},
			want: NewNoneOption[uint](),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Find(testCase.args.values, testCase.args.predicate)
			assert.Equal(t, testCase.want, result)
		})
	}
}
