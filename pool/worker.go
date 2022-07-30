// Package pool defines worker's behavior
// and provide a way to manage multi-goroutine
package pool

import (
	"log"

	"github.com/sinomoe/goworker/work"
)

// a worker represents a goroutine
type worker struct {
	// id is a worker's unique attribute
	id int
	// channel is used to receive new works
	channel chan work.Workable
	// workerChannel holds all available worker's channel
	workerChannel chan chan work.Workable
	// end is used to receive end signal
	end chan bool
}

// start spawn a new goroutine which represents a worker.
// a worker is waiting for 2 channels
// 1. end signal
// 2. works to be done
// a worker sends its channel to workerChannel when having nothing to do
func (w *worker) start() {
	go func() {
		for {
			w.workerChannel <- w.channel
			select {
			case <-w.end:
				return
			case work1 := <-w.channel:
				work1.Do(w.id)
			}
		}
	}()
}

// stop sends end signal to the specified worker
func (w *worker) stop() {
	log.Printf("stopping worker[%d]\n", w.id)
	w.end <- true
}
