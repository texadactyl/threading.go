package helpers

import (
	"sync"
)
import "fmt"

var lock sync.Mutex

func Logger(format string, args ...interface{}) {
	lock.Lock()
	defer lock.Unlock()
	fmt.Printf(format, args...)
}
