package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	type args struct {
		values    []uint
		def       uint
		predicate func(uint, uint) uint
		threads   uint
	}

	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "OK",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				def:       0,
				predicate: func(u1, u2 uint) uint { return u1 + u2 },
			},
			want: 45,
		},
		{
			name: "OK - multihreads",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				def:       0,
				predicate: func(u1, u2 uint) uint { return u1 + u2 },
				threads:   3,
			},
			want: 45,
		},
		{
			name: "OK - with default",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				def:       5,
				predicate: func(u1, u2 uint) uint { return u1 + u2 },
			},
			want: 50,
		},
		{
			name: "OK - with default multithreads",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				def:       5,
				predicate: func(u1, u2 uint) uint { return u1 + u2 },
				threads:   3,
			},
			want: 50,
		},
		{
			name: "no values",
			args: args{
				values:    nil,
				def:       5,
				predicate: func(u1, u2 uint) uint { return u1 + u2 },
			},
			want: 5,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			IteragoThreads = testCase.args.threads
			result := Reduce(testCase.args.values, testCase.args.def, testCase.args.predicate)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleReduce() {
	values := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(Reduce(values, 0, func(acc, value uint8) uint8 { return acc + value }))
	// Output: 45
}
