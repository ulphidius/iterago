package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	type args struct {
		values    []uint
		predicate func(uint, uint) bool
	}

	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "OK - ASC",
			args: args{
				values:    []uint{3, 1, 8, 5, 2, 4, 9, 6, 7, 0},
				predicate: func(u1, u2 uint) bool { return u1 > u2 },
			},
			want: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "OK - DESC",
			args: args{
				values:    []uint{3, 1, 8, 5, 2, 4, 9, 6, 7, 0},
				predicate: func(u1, u2 uint) bool { return u1 < u2 },
			},
			want: []uint{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			name: "OK - already sorted",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u1, u2 uint) bool { return u1 > u2 },
			},
			want: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "OK - single value",
			args: args{
				values:    []uint{3},
				predicate: func(u1, u2 uint) bool { return u1 > u2 },
			},
			want: []uint{3},
		},
		{
			name: "OK - two values",
			args: args{
				values:    []uint{3, 1},
				predicate: func(u1, u2 uint) bool { return u1 > u2 },
			},
			want: []uint{1, 3},
		},
		{
			name: "no values",
			args: args{
				values:    nil,
				predicate: func(u1, u2 uint) bool { return u1 > u2 },
			},
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Run(testCase.name, func(t *testing.T) {
				result := Sort(testCase.args.values, testCase.args.predicate)
				assert.Equal(t, testCase.want, result)
			})
		})
	}
}

func ExampleSort() {
	type student struct {
		Name string
		Age  uint
	}

	students := []student{
		{
			Name: "Max",
			Age:  25,
		},
		{
			Name: "Julie",
			Age:  20,
		},
		{
			Name: "Billy",
			Age:  18,
		},
		{
			Name: "Jacques",
			Age:  60,
		},
		{
			Name: "Ben",
			Age:  10,
		},
	}

	result := Sort(students, func(a student, b student) bool { return a.Age > b.Age })
	result2 := Sort(students, func(a student, b student) bool { return a.Age < b.Age })
	fmt.Println(result, result2)
	// Output: [{Ben 10} {Billy 18} {Julie 20} {Max 25} {Jacques 60}] [{Jacques 60} {Max 25} {Julie 20} {Billy 18} {Ben 10}]
}
