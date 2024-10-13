package main

import (
	"fmt"
	"os"
	"sync"
	"threading/helpers"
	"time"
)

type msgType struct {
	childId  int
	fname    string
	fhandle  *os.File
	nrecords int
}

func main() {

	// Number of child threads.
	NTHREADS := 10
	NRECORDS := 10

	// Threading data.
	var wg sync.WaitGroup
	p2c := make(chan msgType) // parent-to-child channel.
	c2p := make(chan msgType) // child-to-parent channel.

	// Get O/S process ID and thread ID.
	processId := os.Getpid()
	threadId := helpers.GetThreadID()
	helpers.Logger("main: pid = %d, thread_id = %d\n", processId, threadId)

	// Start the clock.
	t1 := time.Now()

	// Start child threads.
	for id := 1; id <= NTHREADS; id++ {
		wg.Add(1)
		go child(p2c, c2p, &wg)
	}

	// Send request messages to the child threads.
	for childId := 1; childId <= NTHREADS; childId++ {
		fname := fmt.Sprintf("%d.%d", processId, childId)
		file, err := os.CreateTemp("", fname)
		if err != nil {
			panic(err)
		}
		p2c <- msgType{
			childId:  childId,
			fname:    fname,
			fhandle:  file,
			nrecords: NRECORDS,
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
		helpers.Logger("main: childId %d completed\n", reply.childId)
	}

	// Stop the clock.
	t2 := time.Now()

	// Report.
	elapsed := t2.Sub(t1)
	helpers.Logger("main: Elapsed time is: %.2f seconds\n", elapsed.Seconds())

}
