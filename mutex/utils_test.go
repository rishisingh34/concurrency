package mutex

import (
    "testing"
)


// TestUpdateMessage appears to pass when run normally (`go test`),
// but fails under the race detector (`go test -race`). This happens because:
//
// 1. Two goroutines write to the shared variable `msg` at the same time.
//    This is a data race, even though the writes are simple string assignments.
//
// 2. Without the `-race` flag, Go does NOT check for race conditions;
//    the test usually ( there is a chance it may FAIL)  "passes" only because one goroutine happens to finish last,
//    making the result look stable.
//
// 3. With `-race` enabled, Go instruments the program and detects the
//    overlapping writes, correctly marking this test as unsafe.
//
// Passing without `-race` does NOT mean the code is correct — it only means
// the race did not get detected at runtime. The code is still racy.

func TestUpdateMessage(t *testing.T) {

	wg.Add(2)
	go updateMessage("hi there")
	go updateMessage("see you again")

	wg.Wait()

	t.Run("see you again test", func (t *testing.T) {
		if msg != "see you again" {
			t.Error("incorrect value in msg")
		}
	})
}