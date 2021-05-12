package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for _, worker := range workers {
		<-worker.done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done
	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}

	go doWork(id, w.in, w.done)
	return w
}

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		done <- true
	}
}

func main() {
	chanDemo()
}
