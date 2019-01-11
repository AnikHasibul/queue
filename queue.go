// package queue gives you a queue group accessibility.
// Helps you to limit goroutines, wait for the end of the all goroutines and much more...
//
//	maxRoutines := 50
//	q := queue.New(maxRoutines)
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

// Q holds a queue group and it's essentials.
type Q struct {
	hasJob chan bool
}

// New creates a new queue group. It takes max running jobs as a parameter.
func New(max int) *Q {
	q := new(Q)
	q.hasJob = make(chan bool, max)
	return q
}

// Add adds a new job to the queue
func (q *Q) Add() {
	q.addJob()
}

// Done decrements the queue group counter.
func (q *Q) Done() {
	q.delJob()
}

// Current returns the current running jobs
func (q *Q) Current() int {
	return len(q.hasJob)
}

// Wait waits for the end of the all jobs.
func (q *Q) Wait() {
	q.waitForEnd()
}

// add jobs till the channel blocks ;)
func (q *Q) addJob() {
	q.hasJob <- true
}

// unblock the channel by receiving from the channel
func (q *Q) delJob() {
	<-q.hasJob
}

// wait until it's 0
func (q *Q) waitForEnd() {
	for len(q.hasJob) != 0 {
	}
}
