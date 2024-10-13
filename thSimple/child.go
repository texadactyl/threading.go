package main

import (
	"os"
	"syscall"
	"threading/helpers"
	"time"
)

// Called as go child(.....)
func child(childId int) {
	defer wg.Done()
	processId := os.Getpid()
	threadId := syscall.Gettid()
	helpers.Logger("child %d: processId = %d, threadId = %d\n", childId, processId, threadId)
	time.Sleep(time.Duration(childId) * time.Second)
	helpers.Logger("child %d: exiting\n", childId)
}
