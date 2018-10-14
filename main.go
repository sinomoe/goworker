package main

import (
	"time"

	"github.com/sino2322/go_worker_pool/pool"
	"github.com/sino2322/go_worker_pool/work"
)

func main() {
	c := pool.StartDispatcher(4)
	works := work.MockSomeWorks(30)

	for i := range works {
		c.Input <- &works[i]
	}
	c.End <- true
	time.Sleep(2 * time.Second)
}
