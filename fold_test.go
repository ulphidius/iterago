package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	type args struct {
		values    []uint
		def       string
		predicate func(string, uint) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				def:       "",
				predicate: func(s string, u uint) string { return fmt.Sprintf("%s%d", s, u) },
			},
			want: "0123456789",
		},
		{
			name: "OK - with default",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				def:       "values: ",
				predicate: func(s string, u uint) string { return fmt.Sprintf("%s%d", s, u) },
			},
			want: "values: 0123456789",
		},
		{
			name: "no values",
			args: args{
				values:    nil,
				def:       "sample",
				predicate: func(s string, u uint) string { return fmt.Sprintf("%s%d", s, u) },
			},
			want: "sample",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Fold(testCase.args.values, testCase.args.def, testCase.args.predicate)
			assert.Equal(t, testCase.want, result)
		})
	}
}
