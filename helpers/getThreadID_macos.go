//go:build darwin
// +build darwin

package helpers

import (
	"syscall"
)

func GetThreadID() uint64 {
	// syscall.RawSyscall invokes the system call to get the thread ID on macOS
	threadID, _, _ := syscall.RawSyscall(syscall.SYS_THREAD_SELFID, 0, 0, 0)
	return uint64(threadID)
}
