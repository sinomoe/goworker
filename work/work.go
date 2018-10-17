// Package work defines Workable interface
// and implement a sample case
package work

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"time"
)

// Work is a sample work type which implemented Workable interface
type Work struct {
	ID  int
	Job string
}

// a workable work must have a Do method
type Workable interface {
	// Do runs the specified work with argument of worker id
	Do(int)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(length int) string {
	rs := make([]rune, length)
	for i := range rs {
		rs[i] = letters[rand.Intn(len(letters))]
	}
	return string(rs)
}

// MockSomeWorks mocks some sample works implemented Workable interface
func MockSomeWorks(amount int) []Work {
	works := make([]Work, amount)
	for i := range works {
		works[i] = Work{i, randomString(8)}
	}
	return works
}

// Do implements Workable interface sample
func (w *Work) Do(workerId int) {
	hash := fnv.New32a()
	hash.Write([]byte(w.Job))
	if os.Getenv("DEBUG") == "true" {
		log.Printf("Worker[%d]: Doing Work[%d] hash word[\"%s\"] to [\"%d\"]\n", workerId, w.ID, w.Job, hash.Sum32())
	}
	time.Sleep(time.Second / 2)
}
