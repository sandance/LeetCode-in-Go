package main

import "sync"

type Task struct {
}

type Pool struct {
	Mu    sync.Mutex
	Tasks []Task
}

func Worker(pool *Pool) {
	for {
		pool.Mu.Lock()
		// begin critical section
		task := pool.Tasks[0]
		pool.Tasks = pool.Tasks[1:]
		pool.Mu.Unlock()
		process(task)
	}
}

func main() {
	pending, done := make(chan *Task), make(chan *Task)
	go sendingWork(pending)

	for i := 0; i < N; i++ {
		go Worker(pending, done)
	}
}

func Worker(in, out chan *Task) {
	input := <-in
	process(t)
	out <- input
}
