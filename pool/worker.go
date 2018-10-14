package pool

import (
	"log"

	"github.com/sinomoe/go_worker_pool/work"
)

type Worker struct {
	ID            int
	WorkerChannel chan chan work.Workable
	Channel       chan work.Workable
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
				work1.DoWork(w.ID)
			}
		}
	}()
}

func (w *Worker) Stop() {
	log.Printf("stopping worker[%d]\n", w.ID)
	w.End <- true
}
