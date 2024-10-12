package main

import "sync"
import "fmt"

var lock sync.Mutex

func logger(format string, args ...interface{}) {
	lock.Lock()
	defer lock.Unlock()
	fmt.Printf(format, args...)
}
