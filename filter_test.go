package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	type args struct {
		values    []uint
		predicate func(uint) bool
	}

	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "OK",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) bool {
					return u%2 == 0
				},
			},
			want: []uint{0, 2, 4, 6, 8},
		},
		{
			name: "no values",
			args: args{
				values: nil,
				predicate: func(u uint) bool {
					return u%2 == 0
				},
			},
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Filter(testCase.args.values, testCase.args.predicate)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleFilter() {
	type user struct {
		Name string
		Age  uint8
	}

	users := []user{
		{
			Name: "Max",
			Age:  15,
		},
		{
			Name: "Michel",
			Age:  25,
		},
		{
			Name: "Julie",
			Age:  19,
		},
		{
			Name: "Sam",
			Age:  35,
		},
	}

	fmt.Println(Filter(users, func(u user) bool { return u.Age > 20 }))
	// Output: [{Michel 25} {Sam 35}]
}
