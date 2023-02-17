package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	type args struct {
		values    []uint
		predicate func(uint) string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "OK",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) string { return fmt.Sprintf("%d", u) },
			},
			want: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
		{
			name: "no values",
			args: args{
				values:    nil,
				predicate: func(u uint) string { return fmt.Sprintf("%d", u) },
			},
			want: nil,
		},
	}

	for _, testCase := range tests {
		result := Map(testCase.args.values, testCase.args.predicate)
		assert.Equal(t, testCase.want, result)
	}
}
