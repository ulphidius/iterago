package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		args []uint
		want []uint
	}{
		{
			name: "OK",
			args: []uint{1, 2, 3, 4, 5},
			want: []uint{5, 4, 3, 2, 1},
		},
		{
			name: "no values",
			args: nil,
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Reverse(testCase.args)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleReverse() {
	values := []uint{1, 2, 3, 4, 5}

	fmt.Println(Reverse(values))
	// Output: [5 4 3 2 1]
}
