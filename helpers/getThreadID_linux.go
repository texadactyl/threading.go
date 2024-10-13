//go:build linux
// +build linux

package helpers

import (
	"syscall"
)

func GetThreadID() uint64 {
	// syscall.RawSyscall invokes the system call to get the thread ID on Linux
	tid, _, _ := syscall.RawSyscall(syscall.SYS_GETTID, 0, 0, 0)
	return uint64(tid)
}
