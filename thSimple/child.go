package main

import (
	"os"
	"syscall"
	"time"
)

// Called as go child(.....)
func child(childId int) {
	defer wg.Done()
	processId := os.Getpid()
	threadId := syscall.Gettid()
	logger("child %d: processId = %d, threadId = %d\n", childId, processId, threadId)
	time.Sleep(time.Duration(childId) * time.Second)
	logger("child %d: exiting\n", childId)
}
