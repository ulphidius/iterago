package iterago

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumerate(t *testing.T) {
	tests := []struct {
		name string
		args []uint
		want []EnumPair[uint]
	}{
		{
			name: "OK",
			args: []uint{1, 2, 3},
			want: []EnumPair[uint]{
				NewEnumPair[uint](0, 1),
				NewEnumPair[uint](1, 2),
				NewEnumPair[uint](2, 3),
			},
		},
		{
			name: "no values",
			args: nil,
			want: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Enumerate(testCase.args)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func ExampleEnumerate() {
	type student struct {
		Name string
		Note uint
	}

	students := []student{
		{
			Name: "Max",
			Note: 50,
		},
		{
			Name: "Julie",
			Note: 25,
		},
		{
			Name: "Sam",
			Note: 75,
		},
	}

	enumStudent := Enumerate(students)

	fmt.Println(enumStudent)
	// Output: [{0 {Max 50}} {1 {Julie 25}} {2 {Sam 75}}]
}
