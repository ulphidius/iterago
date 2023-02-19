package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSorted(t *testing.T) {
	type args struct {
		values    []uint
		predicate func(uint, uint) bool
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "OK - ASC",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u1, u2 uint) bool { return u1 <= u2 },
			},
			want: true,
		},
		{
			name: "OK - DESC",
			args: args{
				values:    []uint{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
				predicate: func(u1, u2 uint) bool { return u1 >= u2 },
			},
			want: true,
		},
		{
			name: "OK - ASC with duplicates",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 4, 5, 6, 7, 8, 8, 9},
				predicate: func(u1, u2 uint) bool { return u1 <= u2 },
			},
			want: true,
		},
		{
			name: "OK - DESC with duplicates",
			args: args{
				values:    []uint{9, 8, 8, 7, 6, 5, 4, 3, 2, 2, 1, 0},
				predicate: func(u1, u2 uint) bool { return u1 >= u2 },
			},
			want: true,
		},
		{
			name: "invalid ASC",
			args: args{
				values:    []uint{0, 1, 2, 3, 10, 5, 6, 7, 8, 9},
				predicate: func(u1, u2 uint) bool { return u1 <= u2 },
			},
			want: false,
		},
		{
			name: "invalid DESC",
			args: args{
				values:    []uint{2, 8, 7, 6, 5, 4, 3, 2, 1, 0},
				predicate: func(u1, u2 uint) bool { return u1 >= u2 },
			},
			want: false,
		},
		{
			name: "single value",
			args: args{
				values:    []uint{1},
				predicate: func(u1, u2 uint) bool { return u1 <= u2 },
			},
			want: true,
		},
		{
			name: "no values",
			args: args{
				values:    nil,
				predicate: func(u1, u2 uint) bool { return u1 <= u2 },
			},
			want: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Run(testCase.name, func(t *testing.T) {
				assert.Equal(t, testCase.want, IsSorted(testCase.args.values, testCase.args.predicate))
			})
		})
	}
}

func ExampleIsSorted() {
	values := []uint{0, 1, 2, 3, 4, 4, 5, 6, 7, 8, 8, 9}
	values2 := []uint{5, 2, 7}
	predicate := func(a uint, b uint) bool { return a <= b }

	fmt.Println(IsSorted(values, predicate), IsSorted(values2, predicate))
	// Output: true false
}
