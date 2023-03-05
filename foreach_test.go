package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForeach(t *testing.T) {
	type args struct {
		values []uint
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
			},
			want: "0123456789",
		},
		{
			name: "no values",
			args: args{},
			want: "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := ""
			Foreach(testCase.args.values, func(value uint) {
				result += fmt.Sprintf("%d", value)
			})
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleForeach() {
	result := ""
	values := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	Foreach(values, func(value uint) { result += fmt.Sprintf("%d", value) })
	fmt.Println(result)
	// Output: 0123456789
}
