package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	type args struct {
		first  []uint
		second []uint
	}
	tests := []struct {
		name string
		args args
		want []Pair[Option[uint]]
	}{
		{
			name: "OK",
			args: args{
				first:  []uint{0, 1, 2, 3, 4},
				second: []uint{5, 6, 7, 8, 9},
			},
			want: []Pair[Option[uint]]{
				NewPair(NewOption[uint](0), NewOption[uint](5)),
				NewPair(NewOption[uint](1), NewOption[uint](6)),
				NewPair(NewOption[uint](2), NewOption[uint](7)),
				NewPair(NewOption[uint](3), NewOption[uint](8)),
				NewPair(NewOption[uint](4), NewOption[uint](9)),
			},
		},
		{
			name: "empty second",
			args: args{
				first: []uint{0, 1, 2, 3, 4},
			},
			want: []Pair[Option[uint]]{
				NewPair(NewOption[uint](0), NewNoneOption[uint]()),
				NewPair(NewOption[uint](1), NewNoneOption[uint]()),
				NewPair(NewOption[uint](2), NewNoneOption[uint]()),
				NewPair(NewOption[uint](3), NewNoneOption[uint]()),
				NewPair(NewOption[uint](4), NewNoneOption[uint]()),
			},
		},
		{
			name: "empty first",
			args: args{
				second: []uint{5, 6, 7, 8, 9},
			},
			want: []Pair[Option[uint]]{
				NewPair(NewNoneOption[uint](), NewOption[uint](5)),
				NewPair(NewNoneOption[uint](), NewOption[uint](6)),
				NewPair(NewNoneOption[uint](), NewOption[uint](7)),
				NewPair(NewNoneOption[uint](), NewOption[uint](8)),
				NewPair(NewNoneOption[uint](), NewOption[uint](9)),
			},
		},
		{
			name: "no values",
			args: args{},
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Zip(testCase.args.first, testCase.args.second)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleZip() {
	first := []uint{0, 1, 2, 3, 4}
	second := []uint{5, 6, 7, 8, 9}

	fmt.Println(Zip(first, second))
	// Output: [{{1 0} {1 5}} {{1 1} {1 6}} {{1 2} {1 7}} {{1 3} {1 8}} {{1 4} {1 9}}]
}
