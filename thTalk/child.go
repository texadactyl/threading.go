package main

import (
	"os"
	"sync"
	"syscall"
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
	threadId := syscall.Gettid()

	var request msgType
	var ok bool
	for {
		request, ok = <-p2c
		if ok {
			break // got a request
		}
	}

	logger("child %d: request cmd=%s, text=%s, processId = %d, threadId = %d\n",
		request.childId, request.cmd, request.text, processId, threadId)
	time.Sleep(time.Duration(request.childId) * time.Second)
	c2p <- msgType{
		childId: request.childId,
		cmd:     request.cmd,
		text:    "exit",
	}
	logger("child %d: exiting\n", request.childId)
}
