package mutex

import (
	"fmt"
	"sync"
)

// SyncMutex shows how using a mutex prevents race conditions when multiple
// goroutines update the same shared variable (`msg`). Only one goroutine can
// enter the critical section at a time, so the final output becomes predictable.
func SyncMutex() {
    msg = "hello world"

    var m sync.Mutex // mutex to protect writes to `msg`

    wg.Add(2)
    go updateMessageTwo("hello universe", &m)
    go updateMessageTwo("hello, cosmos", &m)

    wg.Wait()

    fmt.Println(msgTwo) // deterministic: whichever goroutine acquires the lock last
}
