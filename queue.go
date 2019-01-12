// package queue gives you a queue group accessibility.
// Helps you to limit goroutines, wait for the end of the all goroutines and much more...
//
//	maxRoutines := 50
//	q := queue.New(maxRoutines)
//	defer q.Close()
//	for i := 0; i != 1000; i++ {
//		q.Add()
//		go func(c int) {
//			defer q.Done()
//			fmt.Println(c)
//		}(i)
//	}
//	//wait for the end of the all jobs
//	q.Wait()
package queue

// Q holds a queue group and it's essentials.
type Q struct {
	max        int
	hasJob     chan bool
	waitSignal chan bool
}

// New creates a new queue group. It takes max running jobs as a parameter.
func New(max int) *Q {
	q := new(Q)
	q.max = max
	q.hasJob = make(chan bool, max)
	q.waitSignal = make(chan bool, max)
	return q
}

// Add adds a new job to the queue.
func (q *Q) Add() {
	q.addJob()
}

// Done decrements the queue group counter.
func (q *Q) Done() {
	q.delJob()
	// if channel buffer reaches to the max.
	// replace this with a new channel
	if len(q.waitSignal) == q.max {
		q.waitSignal = make(chan bool, q.max)
	}
	q.waitSignal <- true
}

// Current returns the number of current running jobs.
func (q *Q) Current() int {
	return len(q.hasJob)
}

// Wait waits for the end of the all jobs.
func (q *Q) Wait() {
	q.waitForEnd()
}

// Close closes a queue group gracefully.
func (q *Q) Close() {
	close(q.hasJob)
	close(q.waitSignal)
	q = nil
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
	for {
		if len(q.hasJob) == 0 {
			return
		}
		<-q.waitSignal
	}
}
