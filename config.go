package iterago

import (
	"os"
	"strconv"
)

// Defined number of thread used Iterago
var IteragoThreads uint = 1

func init() {
	threads, err := strconv.Atoi(os.Getenv("ITERAGO_THREADS"))
	if err != nil {
		return
	}

	if threads > int(IteragoThreads) {
		IteragoThreads = uint(threads)
	}
}
