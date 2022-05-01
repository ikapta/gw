package test_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func do(mill time.Duration, wg *sync.WaitGroup) {
  wg.Add(1)

  go func() {
      duration := mill * time.Millisecond
      time.Sleep(duration)
      fmt.Println("后台执行，duration：", duration)
      wg.Done()
  }()
}

func TestWg(t *testing.T) {
  var wg sync.WaitGroup

    go do(100, &wg)
    go do(110, &wg)
    go do(120, &wg)
    go do(130, &wg)

  wg.Wait()
  time.Sleep(1000)
  fmt.Println("Done")
}


// https://stackoverflow.com/questions/70046636/how-to-check-if-sync-waitgroup-done-is-called-in-unit-test