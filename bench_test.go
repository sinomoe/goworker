package main

import (
	"testing"

	"github.com/sinomoe/go_worker_pool/pool"
	"github.com/sinomoe/go_worker_pool/work"
)

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(4)
	for _, w := range work.MockSomeWorks(30) {
		collector.Send(&w)
	}
	collector.End()
}

func BenchmarkSequential(b *testing.B) {
	for _, w := range work.MockSomeWorks(30) {
		w.Do(0)
	}
}
