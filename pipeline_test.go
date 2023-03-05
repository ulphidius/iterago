package iterago

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterMap(t *testing.T) {
	type args struct {
		values []uint
		filter func(uint) bool
		mapper func(uint) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "OK",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				filter: func(u uint) bool { return u%2 == 0 },
				mapper: func(u uint) string { return fmt.Sprintf("%d", u) },
			},
			want: []string{"0", "2", "4", "6", "8"},
		},
		{
			name: "no values",
			args: args{
				filter: func(u uint) bool { return u%2 == 0 },
				mapper: func(u uint) string { return fmt.Sprintf("%d", u) },
			},
			want: nil,
		},
		{
			name: "all filtered",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				filter: func(u uint) bool { return u > 10 },
				mapper: func(u uint) string { return fmt.Sprintf("%d", u) },
			},
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := FilterMap(testCase.args.values, FilterMapPredicates[uint, string]{
				Filter: testCase.args.filter,
				Map:    testCase.args.mapper,
			})
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleFilterMap() {
	values := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	filter := func(value uint) bool { return value%2 == 0 }
	mapper := func(value uint) float64 { return float64(value) / 0.5 }

	result := FilterMap(values, FilterMapPredicates[uint, float64]{
		Filter: filter,
		Map:    mapper,
	})
	fmt.Println(result)
	// Output: [0 4 8 12 16]
}

func TestFilterReduce(t *testing.T) {
	type args struct {
		values []uint
		acc    uint
		filter func(uint) bool
		reduce func(uint, uint) uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "OK",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				acc:    0,
				filter: func(v uint) bool { return v%2 == 0 },
				reduce: func(u1, u2 uint) uint { return u1 + u2 },
			},
			want: 20,
		},
		{
			name: "no values",
			args: args{
				acc:    10,
				filter: func(u uint) bool { return u%2 == 0 },
				reduce: func(u1, u2 uint) uint { return u1 + u2 },
			},
			want: 10,
		},
		{
			name: "all filtered",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				acc:    1,
				filter: func(u uint) bool { return u > 10 },
				reduce: func(u1, u2 uint) uint { return u1 + u2 },
			},
			want: 1,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := FilterReduce(testCase.args.values, testCase.args.acc, FilterReducePredicates[uint]{
				Filter: testCase.args.filter,
				Reduce: testCase.args.reduce,
			})
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleFilterReduce() {
	values := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	filter := func(value uint) bool { return value%2 == 0 }
	reduce := func(acc, value uint) uint { return acc + value }

	result := FilterReduce(values, 0, FilterReducePredicates[uint]{
		Filter: filter,
		Reduce: reduce,
	})
	fmt.Println(result)
	// Output: 20
}

func TestFilterFold(t *testing.T) {
	type args struct {
		values []uint
		acc    string
		filter func(uint) bool
		fold   func(string, uint) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				acc:    "0",
				filter: func(v uint) bool { return v%2 == 0 },
				fold:   func(u1 string, u2 uint) string { return u1 + fmt.Sprintf(",%d", u2) },
			},
			want: "0,0,2,4,6,8",
		},
		{
			name: "no values",
			args: args{
				acc:    "10",
				filter: func(u uint) bool { return u%2 == 0 },
				fold:   func(u1 string, u2 uint) string { return u1 + fmt.Sprintf("%d", u2) },
			},
			want: "10",
		},
		{
			name: "all filtered",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				acc:    "1",
				filter: func(u uint) bool { return u > 10 },
				fold:   func(u1 string, u2 uint) string { return u1 + fmt.Sprintf("%d", u2) },
			},
			want: "1",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := FilterFold(testCase.args.values, testCase.args.acc, FilterFoldPredicates[uint, string]{
				Filter: testCase.args.filter,
				Fold:   testCase.args.fold,
			})
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleFilterFold() {
	values := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	filter := func(value uint) bool { return value%2 == 0 }
	fold := func(acc string, value uint) string { return acc + fmt.Sprintf("%d", value) }

	result := FilterFold(values, "", FilterFoldPredicates[uint, string]{
		Filter: filter,
		Fold:   fold,
	})
	fmt.Println(result)
	// Output: 02468
}

func TestMapReduce(t *testing.T) {
	type args struct {
		values []string
		acc    uint
		mapper func(string) uint
		reduce func(uint, uint) uint
	}

	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "OK",
			args: args{
				values: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
				acc:    0,
				mapper: func(s string) uint {
					result, _ := strconv.Atoi(s)
					return uint(result)
				},
				reduce: func(u1, u2 uint) uint { return u1 + u2 },
			},
			want: 45,
		},
		{
			name: "no values",
			args: args{
				acc: 10,
				mapper: func(s string) uint {
					result, _ := strconv.Atoi(s)
					return uint(result)
				},
				reduce: func(u1, u2 uint) uint { return u1 + u2 },
			},
			want: 10,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := MapReduce(testCase.args.values, testCase.args.acc, MapReducePredicates[string, uint]{
				Map:    testCase.args.mapper,
				Reduce: testCase.args.reduce,
			})
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleMapReduce() {
	values := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	mapper := func(value string) uint { result, _ := strconv.Atoi(value); return uint(result) }
	reduce := func(acc, value uint) uint { return acc + value }

	result := MapReduce(values, 5, MapReducePredicates[string, uint]{
		Map:    mapper,
		Reduce: reduce,
	})
	fmt.Println(result)
	// Output: 50
}

func TestPartitionForeach(t *testing.T) {
	type args struct {
		values []uint
		filter func(uint) bool
	}
	type want struct {
		validate   string
		invalidate string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "OK",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				filter: func(u uint) bool { return u%2 == 0 },
			},
			want: want{
				validate:   "validated values: 02468",
				invalidate: "invalidated values: 13579",
			},
		},
		{
			name: "no values",
			args: args{
				filter: func(u uint) bool { return u%2 == 0 },
			},
			want: want{
				validate:   "validated values: ",
				invalidate: "invalidated values: ",
			},
		},
		{
			name: "all validated",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				filter: func(u uint) bool { return u >= 0 },
			},
			want: want{
				validate:   "validated values: 0123456789",
				invalidate: "invalidated values: ",
			},
		},
		{
			name: "all invalidated",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				filter: func(u uint) bool { return u < 0 },
			},
			want: want{
				validate:   "validated values: ",
				invalidate: "invalidated values: 0123456789",
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			validate := "validated values: "
			invalidates := "invalidated values: "

			PartitionForeach(testCase.args.values, PartitionForeachPredicates[uint]{
				Filter: testCase.args.filter,
				Validate: func(u uint) {
					validate += fmt.Sprintf("%d", u)
				},
				Invalidates: func(u uint) {
					invalidates += fmt.Sprintf("%d", u)
				},
			})

			assert.Equal(t, testCase.want.validate, validate)
			assert.Equal(t, testCase.want.invalidate, invalidates)
		})
	}
}

func ExamplePartitionForeach() {
	values := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	validateResult := "validated values: "
	invalidated := "invalidated values: "

	filter := func(u uint) bool { return u%2 == 0 }
	validate := func(u uint) { validateResult += fmt.Sprintf("%d", u) }
	invalidate := func(u uint) { invalidated += fmt.Sprintf("%d", u) }

	PartitionForeach(values, PartitionForeachPredicates[uint]{
		Filter:      filter,
		Validate:    validate,
		Invalidates: invalidate,
	})

	fmt.Println(validateResult + "\n" + invalidated)
	// Output: validated values: 02468
	// invalidated values: 13579
}
