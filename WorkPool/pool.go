package WorkPool

import (
	"sync"
)

type Worker struct {
	taskChan chan func()
	wg       sync.WaitGroup
}

func NewWorker() *Worker {
	w := &Worker{
		taskChan: make(chan func()),
	}
	w.start()
	return w
}

func (w *Worker) start() {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for task := range w.taskChan {
			task()
		}
	}()
}

func (w *Worker) Stop() {
	close(w.taskChan)
	w.wg.Wait()
}

func (w *Worker) Execute(task func()) {
	w.taskChan <- task
}

func Demo () {
	poolSize := 5
	taskCount := 10

	// 创建一组Worker作为goroutine池
	pool := make([]*Worker, poolSize)
	for i := 0; i < poolSize; i++ {
		pool[i] = NewWorker()
	}

	// 执行一些任务
	for i := 0; i < taskCount; i++ {
		taskID := i
		pool[i%poolSize].Execute(func() {
			// 执行任务的逻辑
			println("Task", taskID)
		})
	}

	// 停止并等待所有Worker完成任务
	for _, w := range pool {
		w.Stop()
	}
}
