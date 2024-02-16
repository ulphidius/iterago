package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOption(t *testing.T) {
	tests := []struct {
		name string
		args uint
		want Option[uint]
	}{
		{
			name: "OK",
			args: 10,
			want: Option[uint]{
				Status: Some,
				Value:  10,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := NewOption(testCase.args)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestNewNoneOption(t *testing.T) {
	tests := []struct {
		name string
		want Option[uint]
	}{
		{
			name: "OK",
			want: Option[uint]{
				Status: None,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := NewNoneOption[uint]()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestOptionIsSome(t *testing.T) {
	tests := []struct {
		name   string
		fields Option[uint]
		want   bool
	}{
		{
			name: "OK",
			fields: Option[uint]{
				Status: Some,
				Value:  1,
			},
			want: true,
		},
		{
			name: "none",
			fields: Option[uint]{
				Status: None,
			},
			want: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.IsSome()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestOptionIsNone(t *testing.T) {
	tests := []struct {
		name   string
		fields Option[uint]
		want   bool
	}{
		{
			name: "OK",
			fields: Option[uint]{
				Status: None,
			},
			want: true,
		},
		{
			name: "some",
			fields: Option[uint]{
				Status: Some,
				Value:  1,
			},
			want: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.IsNone()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestOptionUnwrap(t *testing.T) {
	tests := []struct {
		name         string
		fields       Option[uint]
		want         uint
		wantErr      bool
		errorMessage string
	}{
		{
			name: "OK",
			fields: Option[uint]{
				Status: Some,
				Value:  10,
			},
			want: 10,
		},
		{
			name: "none",
			fields: Option[uint]{
				Status: None,
			},
			wantErr:      true,
			errorMessage: ErrUnwrapNoneOption,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.wantErr {
				assert.PanicsWithValue(t, testCase.errorMessage, func() { testCase.fields.Unwrap() })
				return
			}

			result := testCase.fields.Unwrap()

			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestNewPair(t *testing.T) {
	type args struct {
		first  uint
		second uint
	}
	tests := []struct {
		name string
		args args
		want Pair[uint, uint]
	}{
		{
			name: "OK",
			args: args{
				first:  1,
				second: 2,
			},
			want: Pair[uint, uint]{
				First:  1,
				Second: 2,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := NewPair(testCase.args.first, testCase.args.second)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestNewPairUnwrap(t *testing.T) {
	type want struct {
		first  uint
		second string
	}

	tests := []struct {
		name string
		args Pair[uint, string]
		want want
	}{
		{
			name: "OK",
			args: Pair[uint, string]{First: 10, Second: "ten"},
			want: want{
				first:  10,
				second: "ten",
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			firstResult, secondResult := testCase.args.Unwrap()
			assert.Equal(t, testCase.want.first, firstResult)
			assert.Equal(t, testCase.want.second, secondResult)
		})
	}
}

func TestNewEnumPair(t *testing.T) {
	type args struct {
		index uint
		value int
	}
	tests := []struct {
		name string
		args args
		want EnumPair[int]
	}{
		{
			name: "OK",
			args: args{
				index: 0,
				value: 10,
			},
			want: EnumPair[int]{
				Index: 0,
				Value: 10,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := NewEnumPair(testCase.args.index, testCase.args.value)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestMapIntoList(t *testing.T) {
	type want struct {
		x []string
		y []int64
	}

	tests := []struct {
		name string
		args map[string]int64
		want want
	}{
		{
			name: "OK",
			args: map[string]int64{
				"i": 1,
				"c": 2,
				"p": 3,
			},
			want: want{
				x: []string{"i", "c", "p"},
				y: []int64{1, 2, 3},
			},
		},
		{
			name: "empty map",
			args: map[string]int64{},
			want: want{},
		},
		{
			name: "nil map",
			args: nil,
			want: want{},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result1, result2 := MapIntoList(testCase.args)
			assert.EqualValues(t, testCase.want.x, result1)
			assert.EqualValues(t, testCase.want.y, result2)
		})
	}
}

func ExampleMapIntoList() {
	values := map[string]int64{
		"i":  1,
		"c":  2,
		"cc": 3,
	}

	// Output: [i c cc] [1 2 3]
	fmt.Println(MapIntoList(values))
}
