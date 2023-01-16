package concrete

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ulphidius/iterago"
)

func TestNext(t *testing.T) {
	tests := []struct {
		name   string
		fields *Iterator[uint]
		want   iterago.Option[*Iterator[uint]]
	}{
		{
			name: "OK",
			fields: &Iterator[uint]{
				current: 0,
				next: iterago.Option[*Iterator[uint]]{
					Status: iterago.Some,
					Value: &Iterator[uint]{
						current: 10,
						next: iterago.Option[*Iterator[uint]]{
							Status: iterago.Some,
							Value: &Iterator[uint]{
								current: 5,
								next: iterago.Option[*Iterator[uint]]{
									Status: iterago.None,
								},
							},
						},
					},
				},
			},
			want: iterago.Option[*Iterator[uint]]{
				Status: iterago.Some,
				Value: &Iterator[uint]{
					current: 10,
					next: iterago.Option[*Iterator[uint]]{
						Status: iterago.Some,
						Value: &Iterator[uint]{
							current: 5,
							next: iterago.Option[*Iterator[uint]]{
								Status: iterago.None,
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
				next: iterago.Option[*Iterator[uint]]{
					Status: iterago.None,
				},
			},
			want: iterago.Option[*Iterator[uint]]{
				Status: iterago.None,
			},
		},
		{
			name:   "nil iterator",
			fields: nil,
			want: iterago.Option[*Iterator[uint]]{
				Status: iterago.None,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.want, testCase.fields.Next())
		})
	}
}

func TestHasNext(t *testing.T) {
	tests := []struct {
		name   string
		fields *Iterator[uint]
		want   bool
	}{
		{
			name: "OK",
			fields: &Iterator[uint]{
				current: 0,
				next: iterago.Option[*Iterator[uint]]{
					Status: iterago.Some,
					Value: &Iterator[uint]{
						current: 10,
						next: iterago.Option[*Iterator[uint]]{
							Status: iterago.None,
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
				next: iterago.Option[*Iterator[uint]]{
					Status: iterago.None,
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
				next: iterago.Option[*Iterator[uint]]{
					Status: iterago.Some,
					Value: &Iterator[uint]{
						current: 2,
						next: iterago.Option[*Iterator[uint]]{
							Status: iterago.Some,
							Value: &Iterator[uint]{
								current: 1,
								next: iterago.Option[*Iterator[uint]]{
									Status: iterago.Some,
									Value: &Iterator[uint]{
										current: 5,
										next: iterago.Option[*Iterator[uint]]{
											Status: iterago.Some,
											Value: &Iterator[uint]{
												current: 10,
												next: iterago.Option[*Iterator[uint]]{
													Status: iterago.None,
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
				next:    iterago.NewNoneOption[*Iterator[uint]](),
			},
		},
		{
			name:         "no value",
			args:         nil,
			wantErr:      true,
			errorMessage: iterago.ErrUnwrapNoneOption,
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
