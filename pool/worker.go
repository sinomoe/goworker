package pool

import (
	"log"

	"github.com/sino2322/go_worker_pool/work"
)

type Worker struct {
	ID            int
	WorkerChannel chan chan work.Work
	Channel       chan work.Work
	End           chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			select {
			case <-w.End:
				return
			case work1 := <-w.Channel:
				work.DoWork(w.ID, work1)
			}
		}
	}()
}

func (w *Worker) Stop() {
	log.Printf("stopping worker[%d]\n", w.ID)
	w.End <- true
}
