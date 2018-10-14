package pool

import (
	"log"

	"github.com/sino2322/go_worker_pool/work"
)

type Collector struct {
	Input chan work.Workable
	End   chan bool
}

func StartDispatcher(workerAmount int) *Collector {
	workerChannel := make(chan chan work.Workable, workerAmount)
	workers := make([]Worker, workerAmount)

	// 初始化所有 worker
	// 每个 worker 有自己的 channel, end
	// 所有 worker 公用一个 workerChannel
	for i := range workers {
		workers[i] = Worker{i, workerChannel, make(chan work.Workable), make(chan bool)}
		log.Printf("worker[%d] starting\n", i)
		workers[i].Start()
	}

	input := make(chan work.Workable)
	end := make(chan bool)
	collector := Collector{input, end}
	// 调度Worker
	go func() {
		for {
			select {
			case <-end:
				for i := range workers {
					workers[i].Stop()
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
