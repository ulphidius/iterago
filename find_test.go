package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	type args struct {
		values    []uint
		predicate func(uint) bool
		threads   uint
	}
	tests := []struct {
		name string
		args args
		want Option[uint]
	}{
		{
			name: "OK",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) bool { return u == 5 },
				threads:   1,
			},
			want: NewOption[uint](5),
		},
		{
			name: "OK - Multithreads",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) bool { return u == 5 },
				threads:   3,
			},
			want: NewOption[uint](5),
		},
		{
			name: "not found",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) bool { return u == 666 },
				threads:   1,
			},
			want: NewNoneOption[uint](),
		},
		{
			name: "not found - Multithreads",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) bool { return u == 666 },
				threads:   3,
			},
			want: NewNoneOption[uint](),
		},
		{
			name: "no values",
			args: args{
				values:    nil,
				predicate: func(u uint) bool { return u == 5 },
				threads:   1,
			},
			want: NewNoneOption[uint](),
		},
		{
			name: "no values - Multithreads",
			args: args{
				values:    nil,
				predicate: func(u uint) bool { return u == 5 },
				threads:   3,
			},
			want: NewNoneOption[uint](),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			IteragoThreads = testCase.args.threads
			result := Find(testCase.args.values, testCase.args.predicate)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleFind() {
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

	result := Find(users, func(u user) bool { return u.Age > 20 })

	fmt.Println(result.IsSome(), result.Unwrap())
	// Output: true {Michel 25}
}
