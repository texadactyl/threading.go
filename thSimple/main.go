package main

import (
	"os"
	"sync"
	"threading/helpers"
	"time"
)

var wg sync.WaitGroup

func main() {

	// Number of child threads.
	NTHREADS := 10

	// Get O/S process ID and thread ID.
	processId := os.Getpid()
	threadId := helpers.GetThreadID()
	helpers.Logger("main: pid = %d, thread_id = %d\n", processId, threadId)

	// Start the clock.
	t1 := time.Now()

	// Start child threads.
	for id := 0; id < NTHREADS; id++ {
		wg.Add(1)
		go child(id + 1)
	}

	// Wait for all child threads to finish.
	wg.Wait()

	// Stop the clock.
	t2 := time.Now()

	// Report.
	elapsed := t2.Sub(t1)
	helpers.Logger("main: Elapsed time is: %.2f seconds\n", elapsed.Seconds())

}
