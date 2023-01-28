package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterNext(t *testing.T) {
	upperThan2 := func(u uint) bool { return u > 2 }

	tests := []struct {
		name   string
		fields *Filter[uint]
		want   Option[*Filter[uint]]
	}{
		{
			name: "OK",
			fields: &Filter[uint]{
				current: 0,
				next: Option[*Filter[uint]]{
					Status: Some,
					Value: &Filter[uint]{
						current: 10,
						next: Option[*Filter[uint]]{
							Status: Some,
							Value: &Filter[uint]{
								current: 5,
								next: Option[*Filter[uint]]{
									Status: None,
								},
								predicate: upperThan2,
							},
						},
						predicate: upperThan2,
					},
				},
				predicate: upperThan2,
			},
			want: Option[*Filter[uint]]{
				Status: Some,
				Value: &Filter[uint]{
					current: 10,
					next: Option[*Filter[uint]]{
						Status: Some,
						Value: &Filter[uint]{
							current: 5,
							next: Option[*Filter[uint]]{
								Status: None,
							},
							predicate: upperThan2,
						},
					},
					predicate: upperThan2,
					validated: true,
				},
			},
		},
		{
			name: "none next",
			fields: &Filter[uint]{
				current: 0,
				next: Option[*Filter[uint]]{
					Status: None,
				},
			},
			want: Option[*Filter[uint]]{
				Status: None,
			},
		},
		{
			name:   "nil iterator",
			fields: nil,
			want: Option[*Filter[uint]]{
				Status: None,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Next()

			if testCase.want.IsNone() {
				assert.Equal(t, testCase.want, result)
				return
			}

			expected, _ := testCase.want.Unwrap()

			for result.IsSome() {
				result2, _ := result.Unwrap()
				assert.True(t, result2.equal(expected))
				result = result2.Next()
				expected, _ = expected.Next().Unwrap()
			}
		})
	}
}

func TestFilterHasNext(t *testing.T) {
	tests := []struct {
		name   string
		fields *Filter[uint]
		want   bool
	}{
		{
			name: "OK",
			fields: &Filter[uint]{
				current: 0,
				next: Option[*Filter[uint]]{
					Status: Some,
					Value: &Filter[uint]{
						current: 10,
						next: Option[*Filter[uint]]{
							Status: None,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "none next value",
			fields: &Filter[uint]{
				current: 0,
				next: Option[*Filter[uint]]{
					Status: None,
				},
			},
			want: false,
		},
		{
			name:   "nil iterator",
			fields: nil,
			want:   false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.want, testCase.fields.HasNext())
		})
	}
}

func TestFilterCollect(t *testing.T) {
	upperThan2 := func(u uint) bool { return u > 2 }

	tests := []struct {
		name   string
		fields *Filter[uint]
		want   []uint
	}{
		{
			name: "OK",
			fields: &Filter[uint]{
				current: 0,
				next: Option[*Filter[uint]]{
					Status: Some,
					Value: &Filter[uint]{
						current: 10,
						next: Option[*Filter[uint]]{
							Status: Some,
							Value: &Filter[uint]{
								current: 5,
								next: Option[*Filter[uint]]{
									Status: None,
								},
								predicate: upperThan2,
								validated: true,
							},
						},
						predicate: upperThan2,
						validated: true,
					},
				},
				predicate: upperThan2,
				validated: false,
			},
			want: []uint{10, 5},
		},
		{
			name: "OK - center invalid value",
			fields: &Filter[uint]{
				current: 10,
				next: Option[*Filter[uint]]{
					Status: Some,
					Value: &Filter[uint]{
						current: 1,
						next: Option[*Filter[uint]]{
							Status: Some,
							Value: &Filter[uint]{
								current: 5,
								next: Option[*Filter[uint]]{
									Status: None,
								},
								predicate: upperThan2,
								validated: true,
							},
						},
						predicate: upperThan2,
						validated: false,
					},
				},
				predicate: upperThan2,
				validated: true,
			},
			want: []uint{10, 5},
		},
		{
			name: "OK - last invalid value",
			fields: &Filter[uint]{
				current: 10,
				next: Option[*Filter[uint]]{
					Status: Some,
					Value: &Filter[uint]{
						current: 5,
						next: Option[*Filter[uint]]{
							Status: Some,
							Value: &Filter[uint]{
								current: 1,
								next: Option[*Filter[uint]]{
									Status: None,
								},
								predicate: upperThan2,
								validated: false,
							},
						},
						predicate: upperThan2,
						validated: true,
					},
				},
				predicate: upperThan2,
				validated: true,
			},
			want: []uint{10, 5},
		},
		{
			name: "OK - single value",
			fields: &Filter[uint]{
				current:   10,
				next:      NewNoneOption[*Filter[uint]](),
				predicate: upperThan2,
				validated: true,
			},
			want: []uint{10},
		},
		{
			name: "no valid value",
			fields: &Filter[uint]{
				current: 0,
				next: Option[*Filter[uint]]{
					Status: Some,
					Value: &Filter[uint]{
						current: 1,
						next: Option[*Filter[uint]]{
							Status: Some,
							Value: &Filter[uint]{
								current: 0,
								next: Option[*Filter[uint]]{
									Status: None,
								},
								predicate: upperThan2,
								validated: false,
							},
						},
						predicate: upperThan2,
						validated: false,
					},
				},
				predicate: upperThan2,
				validated: false,
			},
			want: nil,
		},
		{
			name:   "nil filter",
			fields: nil,
			want:   nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Collect()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestFilterFilter(t *testing.T) {
	isUpperThan2 := func(x uint) bool { return x > 1 }
	isLowerThan10 := func(x uint) bool { return x < 10 }
	tests := []struct {
		name   string
		fields *Filter[uint]
		args   func(uint) bool
		want   *Filter[uint]
	}{
		{
			name: "OK",
			fields: &Filter[uint]{
				current:   10,
				predicate: isUpperThan2,
				validated: true,
				next: Option[*Filter[uint]]{
					Status: Some,
					Value: &Filter[uint]{
						current:   1,
						predicate: isUpperThan2,
						validated: false,
						next: Option[*Filter[uint]]{
							Status: Some,
							Value: &Filter[uint]{
								current:   2,
								predicate: isUpperThan2,
								validated: true,
								next: Option[*Filter[uint]]{
									Status: Some,
									Value: &Filter[uint]{
										current:   5,
										predicate: isUpperThan2,
										next: Option[*Filter[uint]]{
											Status: None,
										},
										validated: true,
									},
								},
							},
						},
					},
				},
			},
			args: isLowerThan10,
			want: &Filter[uint]{
				current:   10,
				predicate: isLowerThan10,
				validated: false,
				next: Option[*Filter[uint]]{
					Status: Some,
					Value: &Filter[uint]{
						current:   2,
						predicate: isUpperThan2,
						validated: true,
						next: Option[*Filter[uint]]{
							Status: Some,
							Value: &Filter[uint]{
								current:   5,
								predicate: isUpperThan2,
								next: Option[*Filter[uint]]{
									Status: None,
								},
								validated: true,
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Filter(testCase.args)
			assert.True(t, testCase.want.equal(result))
		})
	}
}

func TestFilterMap(t *testing.T) {
	toString := func(x uint) any { return fmt.Sprintf("%d", x) }
	isUpperThan2 := func(x uint) bool { return x > 2 }
	tests := []struct {
		name   string
		fields *Filter[uint]
		args   func(uint) any
		want   Option[*Mapper[uint, any]]
	}{
		{
			name: "OK",
			args: toString,
			fields: &Filter[uint]{
				current:   10,
				predicate: isUpperThan2,
				validated: true,
				next: Option[*Filter[uint]]{
					Status: Some,
					Value: &Filter[uint]{
						current:   1,
						predicate: isUpperThan2,
						validated: false,
						next: Option[*Filter[uint]]{
							Status: Some,
							Value: &Filter[uint]{
								current:   2,
								predicate: isUpperThan2,
								validated: true,
								next: Option[*Filter[uint]]{
									Status: None,
								},
							},
						},
					},
				},
			},
			want: Option[*Mapper[uint, any]]{
				Status: Some,
				Value: &Mapper[uint, any]{
					current:   10,
					transform: "10",
					predicate: toString,
					next: Option[*Mapper[uint, any]]{
						Status: Some,
						Value: &Mapper[uint, any]{
							current:   8,
							transform: "8",
							predicate: toString,
							next: Option[*Mapper[uint, any]]{
								Status: None,
							},
						},
					},
				},
			},
		},
		{
			name: "OK - without child",
			args: toString,
			fields: &Filter[uint]{
				current:   10,
				predicate: isUpperThan2,
				validated: true,
				next: Option[*Filter[uint]]{
					Status: None,
				},
			},
			want: Option[*Mapper[uint, any]]{
				Status: Some,
				Value: &Mapper[uint, any]{
					current:   10,
					transform: "10",
					predicate: toString,
					next: Option[*Mapper[uint, any]]{
						Status: None,
					},
				},
			},
		},
		{
			name: "OK - without child invalid value",
			args: toString,
			fields: &Filter[uint]{
				current:   1,
				predicate: isUpperThan2,
				validated: false,
				next: Option[*Filter[uint]]{
					Status: None,
				},
			},
			want: Option[*Mapper[uint, any]]{
				Status: None,
			},
		},
		{
			name: "nil filtered",
			args: toString,
			want: Option[*Mapper[uint, any]]{
				Status: None,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Map(testCase.args)
			expected, _ := testCase.want.Unwrap()
			assert.True(t, expected.equal(result))
		})
	}
}

func TestFilteredReduce(t *testing.T) {
	addFunction := func(x uint, y uint) uint { return x + y }
	isUpperThan2 := func(x uint) bool { return x > 2 }

	type args struct {
		accumulator uint
		predicate   func(x uint, y uint) uint
	}
	tests := []struct {
		name   string
		fields *Filter[uint]
		args   args
		want   uint
	}{
		{
			name: "OK",
			args: args{
				accumulator: 0,
				predicate:   addFunction,
			},
			fields: &Filter[uint]{
				current:   10,
				predicate: isUpperThan2,
				validated: true,
				next: Option[*Filter[uint]]{
					Status: Some,
					Value: &Filter[uint]{
						current:   1,
						predicate: isUpperThan2,
						validated: false,
						next: Option[*Filter[uint]]{
							Status: Some,
							Value: &Filter[uint]{
								current:   2,
								predicate: isUpperThan2,
								validated: true,
								next: Option[*Filter[uint]]{
									Status: None,
								},
							},
						},
					},
				},
			},
			want: 12,
		},
		{
			name: "OK - single value",
			args: args{
				accumulator: 5,
				predicate:   addFunction,
			},
			fields: &Filter[uint]{
				current:   10,
				predicate: isUpperThan2,
				validated: true,
				next: Option[*Filter[uint]]{
					Status: None,
				},
			},
			want: 15,
		},
		{
			name: "nil filtered",
			args: args{
				accumulator: 5,
				predicate:   addFunction,
			},
			want: 5,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Reduce(testCase.args.accumulator, testCase.args.predicate)
			assert.Equal(t, testCase.want, result)
		})
	}
}
