package concrete

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ulphidius/iterago/interfaces"
)

func TestNewMapperItem(t *testing.T) {
	toString := func(x uint) string { return fmt.Sprintf("%d", x) }
	type args struct {
		value     uint
		next      interfaces.Option[*Mapper[uint, string]]
		predicate func(uint) string
	}

	tests := []struct {
		name string
		args args
		want *Mapper[uint, string]
	}{
		{
			name: "OK - with child",
			args: args{
				value: 10,
				next: interfaces.NewOption(
					&Mapper[uint, string]{
						current:   150,
						transform: "150",
						next:      interfaces.NewNoneOption[*Mapper[uint, string]](),
						predicate: toString,
					},
				),
				predicate: toString,
			},
			want: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next: interfaces.NewOption(
					&Mapper[uint, string]{
						current:   150,
						transform: "150",
						predicate: toString,
						next:      interfaces.NewNoneOption[*Mapper[uint, string]](),
					},
				),
			},
		},
		{
			name: "OK",
			args: args{
				value:     10,
				next:      interfaces.NewNoneOption[*Mapper[uint, string]](),
				predicate: toString,
			},
			want: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next:      interfaces.NewNoneOption[*Mapper[uint, string]](),
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := NewMapperItem(
				testCase.args.value,
				testCase.args.next,
				testCase.args.predicate,
			)
			assert.True(t, testCase.want.equal(result))
		})
	}
}

func TestMapperCompute(t *testing.T) {
	toString := func(x uint) string { return fmt.Sprintf("%d", x) }
	tests := []struct {
		name   string
		fields *Mapper[uint, string]
		want   interfaces.Option[*Mapper[uint, string]]
	}{
		{
			name: "OK",
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
			},
			want: interfaces.NewOption(
				&Mapper[uint, string]{
					current:   10,
					transform: "10",
					predicate: toString,
				},
			),
		},
		{
			name: "nil mapper",
			want: interfaces.NewNoneOption[*Mapper[uint, string]](),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result, _ := testCase.fields.compute().Unwrap()
			expected, _ := testCase.want.Unwrap()
			assert.True(t, expected.equal(result))
		})
	}
}

func TestMapperHasNext(t *testing.T) {
	tests := []struct {
		name   string
		fields *Mapper[uint, string]
		want   bool
	}{
		{
			name: "OK",
			fields: &Mapper[uint, string]{
				next: interfaces.NewOption(
					&Mapper[uint, string]{},
				),
			},
			want: true,
		},
		{
			name: "none next",
			fields: &Mapper[uint, string]{
				next: interfaces.NewNoneOption[*Mapper[uint, string]](),
			},
			want: false,
		},
		{
			name: "nil next",
			want: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.want, testCase.fields.HasNext())
		})
	}
}

func TestMapperNext(t *testing.T) {
	toString := func(x uint) string { return fmt.Sprintf("%d", x) }
	tests := []struct {
		name   string
		fields *Mapper[uint, string]
		want   interfaces.Option[*Mapper[uint, string]]
	}{
		{
			name: "OK",
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next: interfaces.NewOption(
					&Mapper[uint, string]{
						current:   150,
						predicate: toString,
						next:      interfaces.NewNoneOption[*Mapper[uint, string]](),
					},
				),
			},
			want: interfaces.NewOption(
				&Mapper[uint, string]{
					current:   150,
					transform: "150",
					predicate: toString,
					next:      interfaces.NewNoneOption[*Mapper[uint, string]](),
				},
			),
		},
		{
			name: "no next",
			fields: &Mapper[uint, string]{
				current:   150,
				transform: "150",
				predicate: toString,
				next:      interfaces.NewNoneOption[*Mapper[uint, string]](),
			},
			want: interfaces.NewNoneOption[*Mapper[uint, string]](),
		},
		{
			name: "nil mapper",
			want: interfaces.NewNoneOption[*Mapper[uint, string]](),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Run(testCase.name, func(t *testing.T) {
				result, _ := testCase.fields.Next().Unwrap()
				expected, _ := testCase.want.Unwrap()
				assert.True(t, expected.equal(result))
			})
		})
	}
}

func TestMapperFilter(t *testing.T) {
	toString := func(x uint) string { return fmt.Sprintf("%d", x) }
	isNotNull := func(x string) bool { return len(x) != 0 }
	tests := []struct {
		name   string
		fields *Mapper[uint, string]
		args   func(string) bool
		want   interfaces.Option[*Filtered[string]]
	}{
		{
			name: "OK",
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next: interfaces.NewOption(
					&Mapper[uint, string]{
						current:   150,
						transform: "150",
						predicate: toString,
						next:      interfaces.NewNoneOption[*Mapper[uint, string]](),
					},
				),
			},
			args: isNotNull,
			want: interfaces.Option[*Filtered[string]]{
				Status: interfaces.Some,
				Value: &Filtered[string]{
					current:    "10",
					validated:  true,
					predicates: []func(x string) bool{isNotNull},
					next: interfaces.Option[*Filtered[string]]{
						Status: interfaces.Some,
						Value: &Filtered[string]{
							current:    "150",
							validated:  true,
							predicates: []func(x string) bool{isNotNull},
							next: interfaces.Option[*Filtered[string]]{
								Status: interfaces.None,
							},
						},
					},
				},
			},
		},
		{
			name: "OK - with one filtered value",
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next: interfaces.NewOption(
					&Mapper[uint, string]{
						transform: "",
						predicate: toString,
						next: interfaces.NewOption(
							&Mapper[uint, string]{
								current:   150,
								transform: "150",
								predicate: toString,
								next:      interfaces.NewNoneOption[*Mapper[uint, string]](),
							},
						),
					},
				),
			},
			args: isNotNull,
			want: interfaces.Option[*Filtered[string]]{
				Status: interfaces.Some,
				Value: &Filtered[string]{
					current:    "10",
					validated:  true,
					predicates: []func(x string) bool{isNotNull},
					next: interfaces.Option[*Filtered[string]]{
						Status: interfaces.Some,
						Value: &Filtered[string]{
							current:    "",
							validated:  false,
							predicates: []func(x string) bool{isNotNull},
							next: interfaces.Option[*Filtered[string]]{
								Status: interfaces.Some,
								Value: &Filtered[string]{
									current:    "150",
									validated:  true,
									predicates: []func(x string) bool{isNotNull},
									next: interfaces.Option[*Filtered[string]]{
										Status: interfaces.None,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "OK - without child",
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next:      interfaces.NewNoneOption[*Mapper[uint, string]](),
			},
			args: isNotNull,
			want: interfaces.Option[*Filtered[string]]{
				Status: interfaces.Some,
				Value: &Filtered[string]{
					current:    "10",
					validated:  true,
					predicates: []func(x string) bool{isNotNull},
					next: interfaces.Option[*Filtered[string]]{
						Status: interfaces.None,
					},
				},
			},
		},
		{
			name: "nil mapper",
			args: isNotNull,
			want: interfaces.NewNoneOption[*Filtered[string]](),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result, _ := testCase.fields.Filter(testCase.args).Unwrap()
			expected, _ := testCase.want.Unwrap()
			assert.True(t, expected.equal(result))
		})
	}
}

func TestMapperMap(t *testing.T) {
	toString := func(x uint) string { return fmt.Sprintf("%d", x) }
	toUint := func(x string) any {
		if value, err := strconv.Atoi(x); err == nil {
			return uint(value)
		} else {
			return 0
		}
	}
	tests := []struct {
		name   string
		fields *Mapper[uint, string]
		args   func(string) any
		want   *Mapper[string, any]
	}{
		{
			name: "OK",
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next: interfaces.Option[*Mapper[uint, string]]{
					Status: interfaces.Some,
					Value: &Mapper[uint, string]{
						current:   20,
						transform: "20",
						predicate: toString,
						next: interfaces.Option[*Mapper[uint, string]]{
							Status: interfaces.None,
						},
					},
				},
			},
			args: toUint,
			want: &Mapper[string, any]{
				current:   "10",
				transform: uint(10),
				predicate: toUint,
				next: interfaces.Option[*Mapper[string, any]]{
					Status: interfaces.Some,
					Value: &Mapper[string, any]{
						current:   "20",
						transform: uint(20),
						predicate: toUint,
						next: interfaces.Option[*Mapper[string, any]]{
							Status: interfaces.None,
						},
					},
				},
			},
		},
		{
			name: "OK - without child",
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next: interfaces.Option[*Mapper[uint, string]]{
					Status: interfaces.None,
				},
			},
			args: toUint,
			want: &Mapper[string, any]{
				current:   "10",
				transform: uint(10),
				predicate: toUint,
				next: interfaces.Option[*Mapper[string, any]]{
					Status: interfaces.None,
				},
			},
		},
		{
			name: "nil mapper",
			args: toUint,
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result, _ := testCase.fields.Map(testCase.args).Unwrap()
			assert.True(t, testCase.want.equal(result))
		})
	}
}

func TestMapperCollect(t *testing.T) {
	toString := func(x uint) string { return fmt.Sprintf("%d", x) }
	tests := []struct {
		name   string
		fields *Mapper[uint, string]
		want   []string
	}{
		{
			name: "OK",
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next: interfaces.Option[*Mapper[uint, string]]{
					Status: interfaces.Some,
					Value: &Mapper[uint, string]{
						current:   20,
						transform: "20",
						predicate: toString,
						next: interfaces.Option[*Mapper[uint, string]]{
							Status: interfaces.None,
						},
					},
				},
			},
			want: []string{
				"10",
				"20",
			},
		},
		{
			name: "OK - without child",
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next: interfaces.Option[*Mapper[uint, string]]{
					Status: interfaces.None,
				},
			},
			want: []string{
				"10",
			},
		},
		{
			name: "nil mapper",
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Collect()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestMapperReduce(t *testing.T) {
	toString := func(x uint) string { return fmt.Sprintf("%d", x) }
	concat := func(x string, y string) string {
		if len(x) == 0 {
			return y
		}

		return fmt.Sprintf("%s,%s", x, y)
	}

	type args struct {
		accumulator string
		predicate   func(x, y string) string
	}
	tests := []struct {
		name   string
		args   args
		fields *Mapper[uint, string]
		want   string
	}{
		{
			name: "OK",
			args: args{
				accumulator: "",
				predicate:   concat,
			},
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next: interfaces.Option[*Mapper[uint, string]]{
					Status: interfaces.Some,
					Value: &Mapper[uint, string]{
						current:   20,
						transform: "20",
						predicate: toString,
						next: interfaces.Option[*Mapper[uint, string]]{
							Status: interfaces.None,
						},
					},
				},
			},
			want: "10,20",
		},
		{
			name: "OK - single value",
			args: args{
				accumulator: "",
				predicate:   concat,
			},
			fields: &Mapper[uint, string]{
				current:   10,
				transform: "10",
				predicate: toString,
				next: interfaces.Option[*Mapper[uint, string]]{
					Status: interfaces.None,
				},
			},
			want: "10",
		},
		{
			name: "nil mapper",
			args: args{
				accumulator: "100",
				predicate:   concat,
			},
			fields: nil,
			want:   "100",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Reduce(testCase.args.accumulator, testCase.args.predicate)
			assert.Equal(t, testCase.want, result)
		})
	}
}
