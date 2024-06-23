package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

var pool *ThreadPool
var poolOnce = &sync.Once{}

type Process func()

type ThreadPool struct {
	size      int
	workQueue chan Process
	wg        *sync.WaitGroup
}

func NewThreadPool() *ThreadPool {
	poolOnce.Do(func() {
		size := os.Getenv("THREAD_SIZE")
		poolSize, err := strconv.Atoi(size)
		if err != nil {
			panic(err)
		}

		p := &ThreadPool{
			size:      poolSize,
			workQueue: make(chan Process),
			wg:        &sync.WaitGroup{},
		}
		p.wg.Add(p.size)

		for i := 0; i < p.size; i++ {
			go func() {
				defer p.wg.Done()
				for process := range p.workQueue {
					process()
				}
			}()
		}

		log.Printf("ThreadPool Size: %v\n", p.size)
		pool = p
	})
	return pool
}

func (pool *ThreadPool) Add(process Process) {
	pool.workQueue <- process
}

func (pool *ThreadPool) Wait() {
	close(pool.workQueue)
	pool.wg.Wait()
}
