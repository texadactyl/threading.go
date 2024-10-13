//go:build windows
// +build windows

package helpers

import (
	"syscall"
)

func GetThreadID() uint64 {
	// syscall.GetCurrentThreadId gets the thread ID on Windows
	return uint64(syscall.GetCurrentThreadId())
}
