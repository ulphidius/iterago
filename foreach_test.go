package iterago

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForeach(t *testing.T) {
	type args struct {
		values []uint
		thread uint
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
				thread: 1,
			},
			want: "0,1,2,3,4,5,6,7,8,9,",
		},
		{
			name: "OK - Multithread",
			args: args{
				values: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				thread: 3,
			},
			want: "0,1,2,3,4,5,6,7,8,9,",
		},
		{
			name: "no values",
			args: args{},
			want: "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			iteragoThreads = testCase.args.thread
			mx := new(sync.Mutex)
			result := ""
			Foreach(testCase.args.values, func(value uint) {
				mx.Lock()
				result += fmt.Sprintf("%d,", value)
				mx.Unlock()
			})

			if iteragoThreads > 1 {
				sortedResult := Sort(strings.Split(result, ","), func(a, b string) bool {
					a1, _ := strconv.Atoi(a)
					b1, _ := strconv.Atoi(b)

					return a1 >= b1
				})

				mergedResult := ""
				for _, s := range sortedResult {
					if len(s) == 0 {
						continue
					}
					mergedResult += s + ","
				}
				assert.Equal(t, testCase.want, mergedResult)
				return
			}

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
