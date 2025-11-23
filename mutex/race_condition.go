package mutex

import "fmt" 

// RaceCondition shows how two goroutines writing to the same variable (`msg`)
// at the same time cause unpredictable results. Since there's no mutex,
// whichever goroutine writes last wins, so the output varies between runs.
func RaceCondition() {
    msg = "hello world"

    wg.Add(2)
    go updateMessage("hello universe")
    go updateMessage("hello, cosmos")

    wg.Wait()
    fmt.Println(msg) // unpredictable output
}