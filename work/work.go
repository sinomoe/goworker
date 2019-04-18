// Package work defines Workable interface
// and implement a sample case
package work

import (
	"crypto/sha1"
	"fmt"
	"io"
	"time"
)

// Workable means a workable work must have a Do method
type Workable interface {
	// Do runs the specified work with argument of worker id
	Do(int)
}

type DefaultWork struct {
	hash       string
	createdAt  time.Time
	startedAt  time.Time
	finishedAt time.Time
	f          func(int, *DefaultWork)
}

func (w *DefaultWork) Do(workerID int) {
	w.startedAt = time.Now()
	w.f(workerID, w)
	w.finishedAt = time.Now()
}

func (w *DefaultWork) Hash() string {
	return w.hash
}

func (w *DefaultWork) CreatedAt() time.Time {
	return w.createdAt
}

func (w *DefaultWork) StartedAt() time.Time {
	return w.startedAt
}

func (w *DefaultWork) FinishedAt() time.Time {
	return w.finishedAt
}

func HandleFunc(f func(workerID int, work *DefaultWork)) *DefaultWork {
	now := time.Now()
	h := sha1.New()
	io.WriteString(h, now.String())
	hash := fmt.Sprintf("%x", h.Sum(nil))

	return &DefaultWork{
		hash:      hash,
		createdAt: time.Now(),
		f:         f,
	}
}
