package iterago

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	type args struct {
		values    []uint
		predicate func(uint) string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "OK",
			args: args{
				values:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				predicate: func(u uint) string { return fmt.Sprintf("%d", u) },
			},
			want: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
		{
			name: "no values",
			args: args{
				values:    nil,
				predicate: func(u uint) string { return fmt.Sprintf("%d", u) },
			},
			want: nil,
		},
	}

	for _, testCase := range tests {
		result := Map(testCase.args.values, testCase.args.predicate)
		assert.Equal(t, testCase.want, result)
	}
}

func ExampleMap() {
	type user struct {
		Name string `json:"name"`
		Age  uint8  `json:"age"`
	}

	users := []user{
		{
			Name: "Max",
			Age:  15,
		},
		{
			Name: "Michel",
			Age:  25,
		},
		{
			Name: "Julie",
			Age:  19,
		},
		{
			Name: "Sam",
			Age:  35,
		},
	}

	fmt.Println(
		Map(users, func(u user) string {
			result, _ := json.Marshal(u)
			return string(result)
		}),
	)
	// Output: [{"name":"Max","age":15} {"name":"Michel","age":25} {"name":"Julie","age":19} {"name":"Sam","age":35}]
}
