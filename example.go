package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"time"

	"goworker/pool"
)

// Work is a sample work type which implemented Workable interface
type Work struct {
	ID  int
	Job string
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

func main() {
	c := pool.StartDispatcher(4)
	works := MockSomeWorks(30)

	for i := range works {
		c.Send(&works[i])
	}
	c.End()
}
