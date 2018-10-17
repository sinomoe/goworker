// Package pool defines worker's behavior
// and provide a way to manage multi-goroutine
package pool

import (
	"log"
	"sync"

	"github.com/sinomoe/go_worker_pool/work"
)

// a Collector have
// 1. a work channel to receive new work
// and then send them to available workers.
// 2. an end channel to receive end signal
// and then send them to all workers.
type Collector struct {
	work chan work.Workable
	end  chan bool
	wg   sync.WaitGroup
}

// Send sends work to the specified worker pool
func (c *Collector) Send(work work.Workable) {
	c.work <- work
}

// End sends end signal to the specified worker pool
// and wait until all workers done
func (c *Collector) End() {
	c.end <- true
	c.wg.Wait()
}

// StartDispatcher starts specified numbered workers
// and starts a goroutine to schedule all workers
// returns a Collector pointer, through which we can send new
// work to workers or stop all.
func StartDispatcher(workerAmount int) *Collector {
	workerChannel := make(chan chan work.Workable, workerAmount)
	workers := make([]worker, workerAmount)

	input := make(chan work.Workable)
	end := make(chan bool)
	collector := Collector{
		work: input,
		end:  end,
	}
	collector.wg.Add(workerAmount)

	// init all workers
	// every worker has a unique channel, end
	// all workers share one workerChannel
	for i := range workers {
		workers[i] = worker{i, make(chan work.Workable), workerChannel, make(chan bool)}
		log.Printf("worker[%d] starting\n", i)
		workers[i].start()
	}

	// schedule Workers
	go func() {
		for {
			select {
			case <-end:
				for i := range workers {
					workers[i].stop()
					collector.wg.Done()
				}
				return
			case work1 := <-input:
				worker := <-workerChannel
				worker <- work1
			}
		}
	}()

	return &collector
}
