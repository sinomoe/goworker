package main

import (
	"github.com/sino2322/go_worker_pool/pool"
	"github.com/sino2322/go_worker_pool/work"
)

func main() {
	c := pool.StartDispatcher(4)
	works := work.MockSomeWorks(30)

	for i := range works {
		c.Send(&works[i])
	}
	c.End()
}
