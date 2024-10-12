package main

import (
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func main() {

	// Number of child threads.
	NTHREADS := 20

	// Get O/S process ID and thread ID.
	processId := os.Getpid()
	threadId := syscall.Gettid()
	logger("main: pid = %d, thread_id = %d\n", processId, threadId)

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
	fmt.Printf("main: Elapsed time is: %.2f seconds\n", elapsed.Seconds())

}
