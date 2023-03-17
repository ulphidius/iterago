package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	type args struct {
		values []uint
		number uint
	}
	tests := []struct {
		name string
		args args
		want [][]uint
	}{
		{
			name: "OK",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				number: 3,
			},
			want: [][]uint{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}},
		},
		{
			name: "no values",
			args: args{
				number: 3,
			},
			want: nil,
		},
		{
			name: "number 0",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				number: 0,
			},
			want: [][]uint{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			name: "number 1",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				number: 1,
			},
			want: [][]uint{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Split(testCase.args.values, testCase.args.number)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleSplit() {
	values := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var numberOfSplit uint = 3

	result := Split(values, numberOfSplit)

	fmt.Println(result)
	// Output: [[0 1 2 3] [4 5 6 7] [8 9]]
}
