package iterago

import (
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
			result, err := testCase.fields.Unwrap()

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
