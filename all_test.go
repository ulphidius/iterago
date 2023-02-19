package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	type args struct {
		values    []uint
		predicate func(uint) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "OK",
			args: args{
				values:    []uint{1, 2, 3, 4},
				predicate: func(u uint) bool { return u > 0 },
			},
			want: true,
		},
		{
			name: "with invalid values",
			args: args{
				values:    []uint{1, 2, 0, 4},
				predicate: func(u uint) bool { return u > 0 },
			},
			want: false,
		},
		{
			name: "no values",
			args: args{
				values:    nil,
				predicate: func(u uint) bool { return u > 0 },
			},
			want: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Run(testCase.name, func(t *testing.T) {
				result := All(testCase.args.values, testCase.args.predicate)
				assert.Equal(t, testCase.want, result)
			})
		})
	}
}

func ExampleAll() {
	values := []uint{1, 2, 3, 4, 5}
	values2 := []uint{1, 0, 2, 0, 10}
	fmt.Println(All(values, func(v uint) bool { return v > 0 }), All(values2, func(v uint) bool { return v > 0 }))
	// Output: true false
}
