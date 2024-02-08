package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	type args struct {
		first   []uint
		second  []uint
		threads uint
	}
	tests := []struct {
		name string
		args args
		want []Pair[Option[uint], Option[uint]]
	}{
		{
			name: "OK",
			args: args{
				first:   []uint{0, 1, 2, 3, 4},
				second:  []uint{5, 6, 7, 8, 9},
				threads: 1,
			},
			want: []Pair[Option[uint], Option[uint]]{
				NewPair(NewOption[uint](0), NewOption[uint](5)),
				NewPair(NewOption[uint](1), NewOption[uint](6)),
				NewPair(NewOption[uint](2), NewOption[uint](7)),
				NewPair(NewOption[uint](3), NewOption[uint](8)),
				NewPair(NewOption[uint](4), NewOption[uint](9)),
			},
		},
		{
			name: "OK - Multithreads",
			args: args{
				first:   []uint{0, 1, 2, 3, 4},
				second:  []uint{5, 6, 7, 8, 9},
				threads: 2,
			},
			want: []Pair[Option[uint], Option[uint]]{
				NewPair(NewOption[uint](0), NewOption[uint](5)),
				NewPair(NewOption[uint](1), NewOption[uint](6)),
				NewPair(NewOption[uint](2), NewOption[uint](7)),
				NewPair(NewOption[uint](3), NewOption[uint](8)),
				NewPair(NewOption[uint](4), NewOption[uint](9)),
			},
		},
		{
			name: "OK - Multithreads diff size",
			args: args{
				first:   []uint{0, 1, 2, 3, 4},
				second:  []uint{5, 6, 7},
				threads: 2,
			},
			want: []Pair[Option[uint], Option[uint]]{
				NewPair(NewOption[uint](0), NewOption[uint](5)),
				NewPair(NewOption[uint](1), NewOption[uint](6)),
				NewPair(NewOption[uint](2), NewOption[uint](7)),
				NewPair(NewOption[uint](3), NewNoneOption[uint]()),
				NewPair(NewOption[uint](4), NewNoneOption[uint]()),
			},
		},
		{
			name: "empty second",
			args: args{
				first:   []uint{0, 1, 2, 3, 4},
				threads: 1,
			},
			want: []Pair[Option[uint], Option[uint]]{
				NewPair(NewOption[uint](0), NewNoneOption[uint]()),
				NewPair(NewOption[uint](1), NewNoneOption[uint]()),
				NewPair(NewOption[uint](2), NewNoneOption[uint]()),
				NewPair(NewOption[uint](3), NewNoneOption[uint]()),
				NewPair(NewOption[uint](4), NewNoneOption[uint]()),
			},
		},
		{
			name: "empty second - Multithreads",
			args: args{
				first:   []uint{0, 1, 2, 3, 4},
				threads: 2,
			},
			want: []Pair[Option[uint], Option[uint]]{
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
				second:  []uint{5, 6, 7, 8, 9},
				threads: 1,
			},
			want: []Pair[Option[uint], Option[uint]]{
				NewPair(NewNoneOption[uint](), NewOption[uint](5)),
				NewPair(NewNoneOption[uint](), NewOption[uint](6)),
				NewPair(NewNoneOption[uint](), NewOption[uint](7)),
				NewPair(NewNoneOption[uint](), NewOption[uint](8)),
				NewPair(NewNoneOption[uint](), NewOption[uint](9)),
			},
		},
		{
			name: "empty first - Multithreads",
			args: args{
				second:  []uint{5, 6, 7, 8, 9},
				threads: 2,
			},
			want: []Pair[Option[uint], Option[uint]]{
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
			IteragoThreads = testCase.args.threads
			result := Zip(testCase.args.first, testCase.args.second)
			result = Sort(result, func(a, b Pair[Option[uint], Option[uint]]) bool {
				if a.First.IsSome() {
					return (a.First.Value >= b.First.Value)
				}
				return (a.Second.Value >= b.Second.Value)
			})
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestMapIntoZip(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want []Pair[Option[string], Option[int]]
	}{
		{
			name: "OK",
			args: map[string]int{
				"first":  1,
				"second": 2,
				"third":  3,
			},
			want: []Pair[Option[string], Option[int]]{
				NewPair(NewOption[string]("first"), NewOption[int](1)),
				NewPair(NewOption[string]("second"), NewOption[int](2)),
				NewPair(NewOption[string]("third"), NewOption[int](3)),
			},
		},
		{
			name: "no values",
			args: map[string]int{},
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Sort(MapIntoZip(testCase.args), func(p1, p2 Pair[Option[string], Option[int]]) bool {
				if p1.Second.IsSome() && p2.Second.IsSome() {
					_, v1 := p1.Unwrap()
					_, v2 := p2.Unwrap()
					return v1.Unwrap() >= v2.Unwrap()
				}

				return false
			})
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

func ExampleZip_differentTypes() {
	first := []uint{0, 1, 2, 3, 4}
	second := []string{"zero", "first", "second", "third", "fourth"}

	fmt.Println(Zip(first, second))
	// Output: [{{1 0} {1 zero}} {{1 1} {1 first}} {{1 2} {1 second}} {{1 3} {1 third}} {{1 4} {1 fourth}}]
}

func ExampleMapIntoZip() {
	values := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	result := MapIntoZip(values)
	result = Sort(result, func(p1, p2 Pair[Option[string], Option[int]]) bool {
		if p1.Second.IsSome() && p2.Second.IsSome() {
			_, v1 := p1.Unwrap()
			_, v2 := p2.Unwrap()
			return v1.Unwrap() >= v2.Unwrap()
		}

		return false
	})
	fmt.Println(result)
	// Output: [{{1 first} {1 1}} {{1 second} {1 2}} {{1 third} {1 3}}]
}
