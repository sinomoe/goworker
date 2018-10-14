package pool

import (
	"log"
	"sync"

	"github.com/sino2322/go_worker_pool/work"
)

type Collector struct {
	work chan work.Workable
	end  chan bool
	wg   sync.WaitGroup
}

func (c *Collector) Send(work work.Workable) {
	c.work <- work
}

func (c *Collector) End() {
	c.end <- true
	c.wg.Wait()
}

func StartDispatcher(workerAmount int) *Collector {
	workerChannel := make(chan chan work.Workable, workerAmount)
	workers := make([]Worker, workerAmount)

	input := make(chan work.Workable)
	end := make(chan bool)
	collector := Collector{
		work: input,
		end:  end,
	}
	collector.wg.Add(workerAmount)

	// 初始化所有 worker
	// 每个 worker 有自己的 channel, end
	// 所有 worker 公用一个 workerChannel
	for i := range workers {
		workers[i] = Worker{i, workerChannel, make(chan work.Workable), make(chan bool)}
		log.Printf("worker[%d] starting\n", i)
		workers[i].Start()
	}

	// 调度Worker
	go func() {
		for {
			select {
			case <-end:
				for i := range workers {
					workers[i].Stop()
					collector.wg.Done()
				}
				return
			case work := <-input:
				worker := <-workerChannel
				worker <- work
			}
		}
	}()

	return &collector
}
