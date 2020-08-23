## goworker

![MIT](https://img.shields.io/badge/license-MIT-blue.svg)

An implementation of worker pool pattern for concurrency control in golang.

## Get started

```golang
import (
    "fmt"
    "time"
    
    "goworker/pool"
    "goworker/work"
)

// task A
func myWorkA(workerID int, work *work.DefaultWork) {
    fmt.Printf("Woker[%d] run work[%s]\n", workerID, work.Hash())
    time.Sleep(time.Second / 4)
}

// task B
func myWorkB(workerID int, work *work.DefaultWork) {
    fmt.Printf("Woker[%d] run work[%s]\n", workerID, work.Hash())
    time.Sleep(time.Second / 2)
}

func main() {
    // start pool 
    c := pool.StartDispatcher(4)
    
    // send task to collector
    c.Send(work.HandleFunc(myWorkA))
    c.Send(work.HandleFunc(myWorkB))
    
    // wait end
    c.End()
}
```

## Further

1. define your work type

    ```golang
    type Work struct {
        ID  int
        // ...
    }
    ```

2. implement its Workable interface

    ```golang
    func (w *Work) Do(workerId int) {
        // ...
    }
    ```

3. start a dispatcher and then returns the collector

    ```golang
    c := pool.StartDispatcher(4)
    ```

4. send works to collector

    ```golang
    c.Send(&Work{
        ID: id,
        // ...
    })
    ```

5. stop collector

    ```golang
    c.End()
    ```

for more, see [example.go](example.go)