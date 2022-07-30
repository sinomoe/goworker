package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/sinomoe/goworker/pool"
	"github.com/sinomoe/goworker/work"
)

func mockSomeWorks(n int) []*work.DefaultWork {
	works := make([]*work.DefaultWork, n)
	for i := range works {
		works[i] = work.HandleFunc(func(workerID int, work *work.DefaultWork) {
			if os.Getenv("DEBUG") == "true" {
				fmt.Printf("Woker[%d] run work[%s]\n", workerID, work.Hash())
			}
			time.Sleep(time.Second / 4)
		})
	}
	return works
}

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(4)
	for _, w := range mockSomeWorks(30) {
		collector.Send(w)
	}
	collector.End()
}

func BenchmarkSequential(b *testing.B) {
	for _, w := range mockSomeWorks(30) {
		w.Do(0)
	}
}
