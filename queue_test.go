package queue

import (
	"testing"
)

func TestAdd(t *testing.T) {
	q := New(10)
	n := 5
	for i := 0; i != n; i++ {
		q.Add()
		go func(c int) {
			return
		}(i)
	}
	if jobs := q.Current(); jobs != n {
		t.Errorf("Expected %d got %d", n, jobs)
		t.Fail()
	}
}
func TestWait(t *testing.T) {
	q := New(10)
	n := 5
	for i := 0; i != n; i++ {
		q.Add()
		go func(c int) {
			defer q.Done()
			return
		}(i)
	}
	// wait for the end of the all jobs
	q.Wait()
	if jobs := q.Current(); jobs != 0 {
		t.Errorf("Expected %d got %d", 0, jobs)
		t.Fail()
	}
}
func TestDone(t *testing.T) {
	q := New(10)
	n := 5
	for i := 0; i != n; i++ {
		q.Add()
		go func(c int) {
			// let all the jobs done
			defer q.Done()
			return
		}(i)
	}
	// wait for the end of the all jobs
	q.Wait()
	if jobs := q.Current(); jobs != 0 {
		t.Errorf("Expected %d got %d", 0, jobs)
		t.Fail()
	}
}

func TestCurrent(t *testing.T) {
	q := New(10)
	n := 5
	for i := 0; i != n; i++ {
		q.Add()
		go func(c int) {
			defer q.Done()
			return
		}(i)
	}
	q.Wait()
	// current should be 0
	if jobs := q.Current(); jobs != 0 {
		t.Errorf("Expected %d got %d", 0, jobs)
		t.Fail()
	}
}
