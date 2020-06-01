package helper

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Process()
}

type WorkerPool struct {
	maxWorkers int
	tasks      chan Job
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	return &WorkerPool{
		maxWorkers: maxWorkers,
		tasks:      make(chan Job, 100),
	}
}

func (w *WorkerPool) Close() {
	close(w.tasks)
}

func (w *WorkerPool) Submit(job Job) {
	w.tasks <- job
}
func (w *WorkerPool) Start() {
	wg := &sync.WaitGroup{}
	wg.Add(w.maxWorkers)
	go func() {
		wg.Wait()
	}()

	for i := 1; i <= w.maxWorkers; i++ {
		go func(id int) {
			defer wg.Done()

			for job := range w.tasks {
				fmt.Println("worker", id, "processing job", i)
				job.Process()
				time.Sleep(time.Second)
			}
		}(i)
	}
}
