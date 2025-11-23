package mutex

import "sync"

var msg string

var msgTwo string
var wg sync.WaitGroup

// updateMessageTwo safely updates the shared variable `msg` using a mutex.
// Only one goroutine can lock `m` at a time, preventing race conditions.
func updateMessageTwo(s string, m *sync.Mutex) {
    defer wg.Done()

    m.Lock()        // enter critical section
    msgTwo = s         // safe write
    m.Unlock()      // exit critical section
}


func updateMessage(s string) {
	defer wg.Done()
	msg = s
}