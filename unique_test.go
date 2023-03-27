package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	type sample struct {
		id    uint
		value uint
	}
	type args struct {
		values    []sample
		predicate func(sample) uint
	}
	tests := []struct {
		name string
		args args
		want []sample
	}{
		{
			name: "OK",
			args: args{
				values: []sample{
					{
						id:    0,
						value: 10,
					},
					{
						id:    1,
						value: 20,
					},
					{
						id:    2,
						value: 30,
					},
					{
						id:    1,
						value: 40,
					},
					{
						id:    4,
						value: 50,
					},
				},
				predicate: func(s sample) uint { return s.id },
			},
			want: []sample{
				{
					id:    0,
					value: 10,
				},
				{
					id:    1,
					value: 20,
				},
				{
					id:    2,
					value: 30,
				},
				{
					id:    4,
					value: 50,
				},
			},
		},
		{
			name: "no values",
			args: args{
				predicate: func(s sample) uint { return s.id },
			},
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Unique(testCase.args.values, testCase.args.predicate)
			result = Sort(result, func(a, b sample) bool { return a.id >= b.id })
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleUnique() {
	type sample struct {
		id    uint
		value uint
	}
	values := []sample{
		{
			id:    0,
			value: 10,
		},
		{
			id:    1,
			value: 20,
		},
		{
			id:    2,
			value: 30,
		},
		{
			id:    1,
			value: 40,
		},
		{
			id:    4,
			value: 50,
		},
	}

	result := Unique(values, func(s sample) uint { return s.id })
	result = Sort(result, func(a, b sample) bool { return a.id >= b.id })
	fmt.Println(result)
	// Output: [{0 10} {1 20} {2 30} {4 50}]
}
