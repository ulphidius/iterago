package concrete

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ulphidius/iterago/interfaces"
)

func TestIteratorNext(t *testing.T) {
	tests := []struct {
		name   string
		fields *Iterator[uint]
		want   interfaces.Option[*Iterator[uint]]
	}{
		{
			name: "OK",
			fields: &Iterator[uint]{
				current: 0,
				next: interfaces.Option[*Iterator[uint]]{
					Status: interfaces.Some,
					Value: &Iterator[uint]{
						current: 10,
						next: interfaces.Option[*Iterator[uint]]{
							Status: interfaces.Some,
							Value: &Iterator[uint]{
								current: 5,
								next: interfaces.Option[*Iterator[uint]]{
									Status: interfaces.None,
								},
							},
						},
					},
				},
			},
			want: interfaces.Option[*Iterator[uint]]{
				Status: interfaces.Some,
				Value: &Iterator[uint]{
					current: 10,
					next: interfaces.Option[*Iterator[uint]]{
						Status: interfaces.Some,
						Value: &Iterator[uint]{
							current: 5,
							next: interfaces.Option[*Iterator[uint]]{
								Status: interfaces.None,
							},
						},
					},
				},
			},
		},
		{
			name: "none next",
			fields: &Iterator[uint]{
				current: 0,
				next: interfaces.Option[*Iterator[uint]]{
					Status: interfaces.None,
				},
			},
			want: interfaces.Option[*Iterator[uint]]{
				Status: interfaces.None,
			},
		},
		{
			name:   "nil iterator",
			fields: nil,
			want: interfaces.Option[*Iterator[uint]]{
				Status: interfaces.None,
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
		fields *Iterator[uint]
		want   bool
	}{
		{
			name: "OK",
			fields: &Iterator[uint]{
				current: 0,
				next: interfaces.Option[*Iterator[uint]]{
					Status: interfaces.Some,
					Value: &Iterator[uint]{
						current: 10,
						next: interfaces.Option[*Iterator[uint]]{
							Status: interfaces.None,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "none next value",
			fields: &Iterator[uint]{
				current: 0,
				next: interfaces.Option[*Iterator[uint]]{
					Status: interfaces.None,
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
		fields *Iterator[uint]
		args   func(x uint) bool
		want   interfaces.Option[*Filtered[uint]]
	}{
		{
			name: "OK",
			fields: &Iterator[uint]{
				current: 0,
				next: interfaces.Option[*Iterator[uint]]{
					Status: interfaces.Some,
					Value: &Iterator[uint]{
						current: 10,
						next: interfaces.Option[*Iterator[uint]]{
							Status: interfaces.Some,
							Value: &Iterator[uint]{
								current: 5,
								next: interfaces.Option[*Iterator[uint]]{
									Status: interfaces.None,
								},
							},
						},
					},
				},
			},
			args: upperThan2,
			want: interfaces.NewOption(&Filtered[uint]{
				current: 0,
				next: interfaces.Option[*Filtered[uint]]{
					Status: interfaces.Some,
					Value: &Filtered[uint]{
						current: 10,
						next: interfaces.Option[*Filtered[uint]]{
							Status: interfaces.Some,
							Value: &Filtered[uint]{
								current: 5,
								next: interfaces.Option[*Filtered[uint]]{
									Status: interfaces.None,
								},
								predicates: []func(uint) bool{upperThan2},
							},
						},
						predicates: []func(uint) bool{upperThan2},
					},
				},
				predicates: []func(uint) bool{upperThan2},
			}),
		},
		{
			name: "none next",
			fields: &Iterator[uint]{
				current: 0,
				next: interfaces.Option[*Iterator[uint]]{
					Status: interfaces.None,
				},
			},
			args: upperThan2,
			want: interfaces.Option[*Filtered[uint]]{
				Status: interfaces.Some,
				Value: &Filtered[uint]{
					current: 0,
					next: interfaces.Option[*Filtered[uint]]{
						Status: interfaces.None,
					},
					predicates: []func(uint) bool{upperThan2},
				},
			},
		},
		{
			name:   "nil iterator",
			fields: nil,
			want: interfaces.Option[*Filtered[uint]]{
				Status: interfaces.None,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Filter(testCase.args)

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

func TestSliceUintIntoIter(t *testing.T) {
	tests := []struct {
		name         string
		args         []uint
		want         *Iterator[uint]
		wantErr      bool
		errorMessage string
	}{
		{
			name: "OK",
			args: []uint{0, 2, 1, 5, 10},
			want: &Iterator[uint]{
				current: 0,
				next: interfaces.Option[*Iterator[uint]]{
					Status: interfaces.Some,
					Value: &Iterator[uint]{
						current: 2,
						next: interfaces.Option[*Iterator[uint]]{
							Status: interfaces.Some,
							Value: &Iterator[uint]{
								current: 1,
								next: interfaces.Option[*Iterator[uint]]{
									Status: interfaces.Some,
									Value: &Iterator[uint]{
										current: 5,
										next: interfaces.Option[*Iterator[uint]]{
											Status: interfaces.Some,
											Value: &Iterator[uint]{
												current: 10,
												next: interfaces.Option[*Iterator[uint]]{
													Status: interfaces.None,
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
			want: &Iterator[uint]{
				current: 0,
				next:    interfaces.NewNoneOption[*Iterator[uint]](),
			},
		},
		{
			name:         "no value",
			args:         nil,
			wantErr:      true,
			errorMessage: interfaces.ErrUnwrapNoneOption,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := SliceUintIntoIter(testCase.args).Unwrap()

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
