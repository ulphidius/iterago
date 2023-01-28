package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIteratorNext(t *testing.T) {
	tests := []struct {
		name   string
		fields *Iter[uint]
		want   Option[*Iter[uint]]
	}{
		{
			name: "OK",
			fields: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
					Status: Some,
					Value: &Iter[uint]{
						current: 10,
						next: Option[*Iter[uint]]{
							Status: Some,
							Value: &Iter[uint]{
								current: 5,
								next: Option[*Iter[uint]]{
									Status: None,
								},
							},
						},
					},
				},
			},
			want: Option[*Iter[uint]]{
				Status: Some,
				Value: &Iter[uint]{
					current: 10,
					next: Option[*Iter[uint]]{
						Status: Some,
						Value: &Iter[uint]{
							current: 5,
							next: Option[*Iter[uint]]{
								Status: None,
							},
						},
					},
				},
			},
		},
		{
			name: "none next",
			fields: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
					Status: None,
				},
			},
			want: Option[*Iter[uint]]{
				Status: None,
			},
		},
		{
			name:   "nil iterator",
			fields: nil,
			want: Option[*Iter[uint]]{
				Status: None,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Next()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestIteratorHasNext(t *testing.T) {
	tests := []struct {
		name   string
		fields *Iter[uint]
		want   bool
	}{
		{
			name: "OK",
			fields: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
					Status: Some,
					Value: &Iter[uint]{
						current: 10,
						next: Option[*Iter[uint]]{
							Status: None,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "none next value",
			fields: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
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

func TestIteratorFilter(t *testing.T) {
	upperThan2 := func(u uint) bool { return u > 2 }

	tests := []struct {
		name   string
		fields *Iter[uint]
		args   func(x uint) bool
		want   *Filter[uint]
	}{
		{
			name: "OK",
			fields: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
					Status: Some,
					Value: &Iter[uint]{
						current: 10,
						next: Option[*Iter[uint]]{
							Status: Some,
							Value: &Iter[uint]{
								current: 5,
								next: Option[*Iter[uint]]{
									Status: None,
								},
							},
						},
					},
				},
			},
			args: upperThan2,
			want: &Filter[uint]{
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
		},
		{
			name: "none next",
			fields: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
					Status: None,
				},
			},
			args: upperThan2,
			want: &Filter[uint]{
				current: 0,
				next: Option[*Filter[uint]]{
					Status: None,
				},
				predicate: upperThan2,
			},
		},
		{
			name:   "nil iterator",
			fields: nil,
			want:   nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Filter(testCase.args)

			for result != nil {
				assert.True(t, result.equal(testCase.want))
				result, _ = result.Next().Unwrap()
				testCase.want, _ = testCase.want.Next().Unwrap()
			}
		})
	}
}

func TestIteractorCollect(t *testing.T) {
	tests := []struct {
		name   string
		fields *Iter[uint]
		want   []uint
	}{
		{
			name: "OK",
			fields: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
					Status: Some,
					Value: &Iter[uint]{
						current: 10,
						next: Option[*Iter[uint]]{
							Status: Some,
							Value: &Iter[uint]{
								current: 5,
								next: Option[*Iter[uint]]{
									Status: None,
								},
							},
						},
					},
				},
			},
			want: []uint{0, 10, 5},
		},
		{
			name: "OK - Single value",
			fields: &Iter[uint]{
				current: 0,
				next:    NewNoneOption[*Iter[uint]](),
			},
			want: []uint{0},
		},
		{
			name:   "nil value",
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

func TestIteractorMap(t *testing.T) {
	toString := func(x uint) any { return fmt.Sprintf("%d", x) }

	tests := []struct {
		name   string
		fields *Iter[uint]
		args   func(x uint) any
		want   *Mapper[uint, any]
	}{
		{
			name: "OK",
			fields: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
					Status: Some,
					Value: &Iter[uint]{
						current: 10,
						next: Option[*Iter[uint]]{
							Status: Some,
							Value: &Iter[uint]{
								current: 5,
								next: Option[*Iter[uint]]{
									Status: None,
								},
							},
						},
					},
				},
			},
			args: toString,
			want: &Mapper[uint, any]{
				current:   0,
				transform: "0",
				predicate: toString,
				next: NewOption(
					&Mapper[uint, any]{
						current:   10,
						transform: "10",
						predicate: toString,
						next: NewOption(
							&Mapper[uint, any]{
								current:   5,
								transform: "5",
								predicate: toString,
								next:      NewNoneOption[*Mapper[uint, any]](),
							},
						),
					},
				),
			},
		},
		{
			name: "OK - Single value",
			fields: &Iter[uint]{
				current: 0,
				next:    NewNoneOption[*Iter[uint]](),
			},
			args: toString,
			want: &Mapper[uint, any]{
				current:   0,
				transform: "0",
				predicate: toString,
				next:      NewNoneOption[*Mapper[uint, any]](),
			},
		},
		{
			name:   "nil value",
			fields: nil,
			want:   nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Map(testCase.args)

			for result != nil {
				assert.True(t, result.equal(testCase.want))
				result, _ = result.Next().Unwrap()
				testCase.want, _ = testCase.want.Next().Unwrap()
			}
		})
	}
}

func TestSliceIntoIter(t *testing.T) {
	tests := []struct {
		name         string
		args         []uint
		want         *Iter[uint]
		wantErr      bool
		errorMessage string
	}{
		{
			name: "OK",
			args: []uint{0, 2, 1, 5, 10},
			want: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
					Status: Some,
					Value: &Iter[uint]{
						current: 2,
						next: Option[*Iter[uint]]{
							Status: Some,
							Value: &Iter[uint]{
								current: 1,
								next: Option[*Iter[uint]]{
									Status: Some,
									Value: &Iter[uint]{
										current: 5,
										next: Option[*Iter[uint]]{
											Status: Some,
											Value: &Iter[uint]{
												current: 10,
												next: Option[*Iter[uint]]{
													Status: None,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "OK - single value",
			args: []uint{0},
			want: &Iter[uint]{
				current: 0,
				next:    NewNoneOption[*Iter[uint]](),
			},
		},
		{
			name:         "no value",
			args:         nil,
			wantErr:      true,
			errorMessage: ErrUnwrapNoneOption,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := SliceIntoIter(testCase.args).Unwrap()

			if testCase.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, testCase.errorMessage, err.Error())
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestIteratorReduce(t *testing.T) {
	addFunction := func(u1, u2 uint) uint { return u1 + u2 }
	type args struct {
		accumulator uint
		predicate   func(uint, uint) uint
	}
	tests := []struct {
		name   string
		fields *Iter[uint]
		args   args
		want   uint
	}{
		{
			name: "OK",
			fields: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
					Status: Some,
					Value: &Iter[uint]{
						current: 2,
						next: Option[*Iter[uint]]{
							Status: Some,
							Value: &Iter[uint]{
								current: 1,
								next: Option[*Iter[uint]]{
									Status: Some,
									Value: &Iter[uint]{
										current: 5,
										next: Option[*Iter[uint]]{
											Status: Some,
											Value: &Iter[uint]{
												current: 10,
												next: Option[*Iter[uint]]{
													Status: None,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			args: args{
				accumulator: 0,
				predicate:   addFunction,
			},
			want: 18,
		},
		{
			name: "OK - single value",
			fields: &Iter[uint]{
				current: 0,
				next: Option[*Iter[uint]]{
					Status: None,
				},
			},
			args: args{
				accumulator: 5,
				predicate:   addFunction,
			},
			want: 5,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Reduce(
				testCase.args.accumulator,
				testCase.args.predicate,
			)
			assert.Equal(t, testCase.want, result)
		})
	}
}
