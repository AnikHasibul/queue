# queue
--
    import "github.com/anikhasibul/queue"

package queue gives you a queue group accessibility. Helps you to limit
goroutines, wait for the end of the all goroutines and much more...

    maxRoutines := 50
    q := New(maxRoutines)
    for i := 0; i != 1000; i++ {
    	q.Add()
    	go func(c int) {
    		defer q.Done()
    		fmt.Println(c)
    		return
    	}(i)
    }
    //wait for the end of the all jobs
    q.Wait()

## Usage

#### type Q

```go
type Q struct {
}
```

Q holds a queue group and it's essentials.

#### func  New

```go
func New(max int) *Q
```
New creates a new queue group. It takes max running jobs as a parameter.

#### func (*Q) Add

```go
func (q *Q) Add()
```
Add adds a new job to the queue

#### func (*Q) Current

```go
func (q *Q) Current() int
```
Current returns the current running jobs

#### func (*Q) Done

```go
func (q *Q) Done()
```
Done decrements the queue group counter.

#### func (*Q) Wait

```go
func (q *Q) Wait()
```
Wait waits for the end of the all jobs.
