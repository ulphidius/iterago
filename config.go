package iterago

import (
	"os"
	"strconv"
)

// Defined number of thread used Iterago
var iteragoThreads uint = 1

func init() {
	threads, err := strconv.Atoi(os.Getenv("ITERAGO_THREADS"))
	if err != nil {
		return
	}

	if threads > int(iteragoThreads) {
		iteragoThreads = uint(threads)
	}
}
