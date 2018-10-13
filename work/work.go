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

func DoWork(workerId int, work Work) {
	hash := fnv.New32a()
	hash.Write([]byte(work.Job))
	if os.Getenv("DEBUG") == "true" {
		log.Printf("Worker[%d]: Doing Work[%d] hash word[\"%s\"] to [\"%d\"]\n", workerId, work.ID, work.Job, hash.Sum32)
	}
	time.Sleep(time.Second / 2)
}
