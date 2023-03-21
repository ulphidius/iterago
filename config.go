package iterago

import (
	"fmt"
	"os"
	"strconv"
)

// Defined number of thread used Iterago
var IteragoThreads uint = 1

func init() {
	fmt.Println("INIT", os.Getenv("ITERAGO_THREADS"))
	threads, err := strconv.Atoi(os.Getenv("ITERAGO_THREADS"))
	fmt.Println(threads)
	if err != nil {
		return
	}

	if threads > int(IteragoThreads) {
		IteragoThreads = uint(threads)
	}
}
