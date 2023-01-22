package concrete

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ulphidius/iterago/interfaces"
)

func TestFilterNext(t *testing.T) {
	upperThan2 := func(u uint) bool { return u > 2 }

	tests := []struct {
		name   string
		fields *Filtered[uint]
		want   interfaces.Option[*Filtered[uint]]
	}{
		{
			name: "OK",
			fields: &Filtered[uint]{
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
			},
			want: interfaces.Option[*Filtered[uint]]{
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
					validated:  true,
				},
			},
		},
		{
			name: "none next",
			fields: &Filtered[uint]{
				current: 0,
				next: interfaces.Option[*Filtered[uint]]{
					Status: interfaces.None,
				},
			},
			want: interfaces.Option[*Filtered[uint]]{
				Status: interfaces.None,
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
		fields *Filtered[uint]
		want   bool
	}{
		{
			name: "OK",
			fields: &Filtered[uint]{
				current: 0,
				next: interfaces.Option[*Filtered[uint]]{
					Status: interfaces.Some,
					Value: &Filtered[uint]{
						current: 10,
						next: interfaces.Option[*Filtered[uint]]{
							Status: interfaces.None,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "none next value",
			fields: &Filtered[uint]{
				current: 0,
				next: interfaces.Option[*Filtered[uint]]{
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

func TestFilterCollect(t *testing.T) {
	upperThan2 := func(u uint) bool { return u > 2 }

	tests := []struct {
		name   string
		fields *Filtered[uint]
		want   []uint
	}{
		{
			name: "OK",
			fields: &Filtered[uint]{
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
								validated:  true,
							},
						},
						predicates: []func(uint) bool{upperThan2},
						validated:  true,
					},
				},
				predicates: []func(uint) bool{upperThan2},
				validated:  false,
			},
			want: []uint{10, 5},
		},
		{
			name: "OK - center invalid value",
			fields: &Filtered[uint]{
				current: 10,
				next: interfaces.Option[*Filtered[uint]]{
					Status: interfaces.Some,
					Value: &Filtered[uint]{
						current: 1,
						next: interfaces.Option[*Filtered[uint]]{
							Status: interfaces.Some,
							Value: &Filtered[uint]{
								current: 5,
								next: interfaces.Option[*Filtered[uint]]{
									Status: interfaces.None,
								},
								predicates: []func(uint) bool{upperThan2},
								validated:  true,
							},
						},
						predicates: []func(uint) bool{upperThan2},
						validated:  false,
					},
				},
				predicates: []func(uint) bool{upperThan2},
				validated:  true,
			},
			want: []uint{10, 5},
		},
		{
			name: "OK - last invalid value",
			fields: &Filtered[uint]{
				current: 10,
				next: interfaces.Option[*Filtered[uint]]{
					Status: interfaces.Some,
					Value: &Filtered[uint]{
						current: 5,
						next: interfaces.Option[*Filtered[uint]]{
							Status: interfaces.Some,
							Value: &Filtered[uint]{
								current: 1,
								next: interfaces.Option[*Filtered[uint]]{
									Status: interfaces.None,
								},
								predicates: []func(uint) bool{upperThan2},
								validated:  false,
							},
						},
						predicates: []func(uint) bool{upperThan2},
						validated:  true,
					},
				},
				predicates: []func(uint) bool{upperThan2},
				validated:  true,
			},
			want: []uint{10, 5},
		},
		{
			name: "OK - single value",
			fields: &Filtered[uint]{
				current:    10,
				next:       interfaces.NewNoneOption[*Filtered[uint]](),
				predicates: []func(uint) bool{upperThan2},
				validated:  true,
			},
			want: []uint{10},
		},
		{
			name: "no valid value",
			fields: &Filtered[uint]{
				current: 0,
				next: interfaces.Option[*Filtered[uint]]{
					Status: interfaces.Some,
					Value: &Filtered[uint]{
						current: 1,
						next: interfaces.Option[*Filtered[uint]]{
							Status: interfaces.Some,
							Value: &Filtered[uint]{
								current: 0,
								next: interfaces.Option[*Filtered[uint]]{
									Status: interfaces.None,
								},
								predicates: []func(uint) bool{upperThan2},
								validated:  false,
							},
						},
						predicates: []func(uint) bool{upperThan2},
						validated:  false,
					},
				},
				predicates: []func(uint) bool{upperThan2},
				validated:  false,
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
