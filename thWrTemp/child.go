package main

import (
	"log"
	"os"
	"sync"
	"threading/helpers"
	"time"
)

// Called as: go child.
// Read a request message from the p2c (parent-to-child) channel.
// Accept child ID.
// Sleep a bit.
// Write a reply message to the c2p (child-to-parent) channel.
func child(p2c chan msgType, c2p chan msgType, wg *sync.WaitGroup) {
	defer wg.Done()
	processId := os.Getpid()
	threadId := helpers.GetThreadID()

	var request msgType
	var ok bool
	var err error

	// Get a request message.
	for {
		request, ok = <-p2c
		if ok {
			break // got a request
		}
	}
	helpers.Logger("child %d: request fname=%s, nrecords=%d, processId = %d, threadId = %d\n",
		request.childId, request.fname, request.nrecords, processId, threadId)

	// Write the requested number of records.
	for ii := 0; ii < request.nrecords; ii++ {
		_, err = request.fhandle.Write([]byte{byte(ii), 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		if err != nil {
			log.Panicf("child %d: write error: %s\n", request.childId, err)
		}
		helpers.Logger("child %d: record #%d written\n", request.childId, ii+1)
		time.Sleep(time.Duration(request.childId*100) * time.Millisecond)
	}
	err = request.fhandle.Close()
	if err != nil {
		log.Panicf("child %d: close error: %s\n", request.childId, err)
	}
	err = os.Remove(request.fhandle.Name())
	if err != nil {
		log.Panicf("child %d: remove error: %s\n", request.childId, err)
	}

	// Send the reply message (same as the request message).
	c2p <- request
	helpers.Logger("child %d: exiting\n", request.childId)
}
