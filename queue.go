// package queue gives you a queue group accessibility.
// Helps you to limit goroutines, wait for the end of the all goroutines and much more...
//
//	maxRoutines := 50
//	q := New(maxRoutines)
//	for i := 0; i != 1000; i++ {
//		q.Add()
//		go func(c int) {
//			defer q.Done()
//			fmt.Println(c)
//			return
//		}(i)
//	}
//	//wait for the end of the all jobs
//	q.Wait()
package queue

import (
	"sync/atomic"
)

// Q holds a queue group and it's essentials.
type Q struct {
	max     int
	running int32
	job     chan int
	done    chan int
	play    chan int
	proc    chan func()
}

// New creates a new queue group. It takes max running jobs as a parameter.
func New(max int) *Q {
	q := new(Q)
	q.max = max
	q.job = make(chan int, 1)
	q.done = make(chan int, 1)
	q.play = make(chan int, 1)
	go q.run()
	return q
}

// Add adds a new job to the queue
func (q *Q) Add() {
	q.job <- 1
	<-q.play
}

// Done decrements the queue group counter.
func (q *Q) Done() {
	q.done <- 1
}

// Current returns the current running jobs
func (q *Q) Current() int {
	return int(atomic.LoadInt32(&q.running))
}

// Wait waits for the end of the all jobs.
func (q *Q) Wait() {
	for {
		if q.Current() != 0 {
			<-q.play
		} else {
			break
		}
	}
}

func (q *Q) run() {
	for {
		select {
		case <-q.job:
			atomic.AddInt32(&q.running, 1)
			if q.Current() <= q.max {
				q.play <- 1
			}
		case <-q.done:
			atomic.AddInt32(&q.running, -1)

			if q.Current() <= q.max {
				q.play <- 1
			}
		}
	}
}
