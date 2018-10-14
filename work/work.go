package work

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"time"
)

type Work struct {
	ID  int
	Job string
}

type Workable interface {
	DoWork(int)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(length int) string {
	rs := make([]rune, length)
	for i := range rs {
		rs[i] = letters[rand.Intn(len(letters))]
	}
	return string(rs)
}

func MockSomeWorks(amount int) []Work {
	works := make([]Work, amount)
	for i := range works {
		works[i] = Work{i, randomString(8)}
	}
	return works
}

func (w *Work) DoWork(workerId int) {
	hash := fnv.New32a()
	hash.Write([]byte(w.Job))
	if os.Getenv("DEBUG") == "true" {
		log.Printf("Worker[%d]: Doing Work[%d] hash word[\"%s\"] to [\"%d\"]\n", workerId, w.ID, w.Job, hash.Sum32())
	}
	time.Sleep(time.Second / 2)
}
