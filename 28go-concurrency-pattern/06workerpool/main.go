package main

import (
	"fmt"
	"time"
)

type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker   int
	taskQueueCh chan func()
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		fmt.Println("I am Worker with ID ", i+1)
		go func(workerId int) {
			for task := range wp.taskQueueCh {
				fmt.Printf("I am worker %d, task Start\n", workerId)
				task()
				fmt.Printf("I am worker %d, task End\n", workerId)
			}
		}(i + 1)
	}
}

func (wp *workerPool) AddTask(task func()) {
	wp.taskQueueCh <- task
}

func NewWorkerPool(maxWorker int) *workerPool {
	wp := &workerPool{
		maxWorker:   maxWorker,
		taskQueueCh: make(chan func()),
	}
	wp.maxWorker = maxWorker
	return wp

}
func main() {
	maxWorker := 3
	wp := NewWorkerPool(maxWorker)
	wp.Run()

	tasks := 5
	for i := 0; i < tasks; i++ {
		wp.AddTask(func() {
			fmt.Println("I am task running with ID", i+1)
			time.Sleep(time.Second * 5)
		})
	}
}
