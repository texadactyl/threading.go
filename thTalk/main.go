package main

import (
	"os"
	"sync"
	"syscall"
	"time"
)

type msgType struct {
	childId int
	cmd     string
	text    string
}

func main() {

	// Number of child threads.
	NTHREADS := 20

	// Threading data.
	var wg sync.WaitGroup
	p2c := make(chan msgType) // parent-to-child channel.
	c2p := make(chan msgType) // child-to-parent channel.

	// Get O/S process ID and thread ID.
	processId := os.Getpid()
	threadId := syscall.Gettid()
	logger("main: pid = %d, thread_id = %d\n", processId, threadId)

	// Start the clock.
	t1 := time.Now()

	// Start child threads.
	for id := 1; id <= NTHREADS; id++ {
		wg.Add(1)
		go child(p2c, c2p, &wg)
	}

	// Send request messages to the child threads.
	for id := 1; id <= NTHREADS; id++ {
		p2c <- msgType{
			childId: id,
			cmd:     "begin",
			text:    "now",
		}
	}

	// Collect replies.
	var reply msgType
	var ok bool
	for id := 1; id <= NTHREADS; id++ {
		for {
			reply, ok = <-c2p
			if ok {
				break // got a reply
			}
		}
		logger("main: childId %d completed\n", reply.childId)
	}

	// Stop the clock.
	t2 := time.Now()

	// Report.
	elapsed := t2.Sub(t1)
	logger("main: Elapsed time is: %.2f seconds\n", elapsed.Seconds())

}
