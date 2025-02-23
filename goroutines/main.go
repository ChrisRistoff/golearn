package main

import (
    "fmt"
    "sync"
    "time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

// note: m is a rw mutex for synchronizing changes to the shared 'results' slice.
var m = sync.RWMutex{}

// note: wg is a wait group to wait until all goroutines have finished.
var wg = sync.WaitGroup{}

var results = []string{}

func main() {
    t0 := time.Now()
    _ = dbCall

    for i := 0; i < 1000; i++ {
        // note: increment the wait group counter before starting a goroutine.
        wg.Add(1)
        go count()
    }

    // note: wait for all goroutines to complete.
    wg.Wait()

    fmt.Println("execution time: ", time.Since(t0))
    fmt.Println("the results are: ", results)
}

func count() {
    var res int

	for i := 0; i < 100000000; i++ {
        res++
    }

    // note: signal that this goroutine is done.
    wg.Done()
}

func dbCall(i int) {
    // note: simulate latency by sleeping for a specified delay.
    var delay float32 = 2000
    time.Sleep(time.Duration(delay) * time.Millisecond)

    // note: save the result from the simulated database call.
    save(dbData[i])

    log()

    // note: signal that this goroutine is done.
    wg.Done()
}

func log() {
    // note: acquire a read lock to safely print the shared 'results' slice.
    m.RLock()
    fmt.Println("current results are: ", results)
    m.RUnlock()
}

func save(result string) {
    // note: acquire a write lock before modifying the shared 'results' slice
    // to avoid race conditions.
    m.Lock()
    results = append(results, result)
    m.Unlock()
}
