//go:build windows
// +build windows

package helpers

import (
	"golang.org/x/sys/windows"
)

func GetThreadID() uint64 {
	// windows.GetCurrentThreadId gets the thread ID on Windows
	return uint64(windows.GetCurrentThreadId())
}
