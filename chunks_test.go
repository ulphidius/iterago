package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunks(t *testing.T) {
	type args struct {
		values []uint
		size   uint
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
				size:   2,
			},
			want: [][]uint{
				{0, 1},
				{2, 3},
				{4, 5},
				{6, 7},
				{8, 9},
			},
		},
		{
			name: "OK - without all sub array filled",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8},
				size:   2,
			},
			want: [][]uint{
				{0, 1},
				{2, 3},
				{4, 5},
				{6, 7},
				{8},
			},
		},
		{
			name: "chunk size 0",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				size:   0,
			},
			want: [][]uint{
				{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name: "chunk bigger than array len",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				size:   666,
			},
			want: [][]uint{
				{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name: "no values",
			args: args{
				values: nil,
				size:   666,
			},
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Chunks(testCase.args.values, testCase.args.size)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleChunks() {
	values := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	values2 := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(Chunks(values, 2), Chunks(values2, 2))
	// Output: [[0 1] [2 3] [4 5] [6 7] [8 9]] [[0 1] [2 3] [4 5] [6 7] [8]]
}
