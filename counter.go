package counter

import "sync"

type incrementable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

type Counter[T any] interface {
	// Increament() increments the counter value by 1. It is go routine safe. Any number of go routines
	// can call this method at the same time.
	Increament()
	// Decreament() decrements the counter value by 1. It is go routine safe. Any number of go routines
	// can call this method at the same time.
	Decreament()
	// GetCount() returns the current value of the counter.
	GetCount() T
	// GetCountAndReset() returns the current value of the counter and resets the counter to 0.
	GetCountAndReset() T
	// IncrementBy(T) increments the counter value by the value passed in. It is go routine safe. Any number of go routines
	// can call this method at the same time.
	IncrementBy(T)
	// DecrementBy(T) decrements the counter value by the value passed in. It is go routine safe. Any number of go routines
	// can call this method at the same time.
	DecrementBy(T)
	// Reset() resets the counter value to 0.
	Reset()
	// SetCount(T) sets the counter value to the value passed in.
	SetCount(T)
	// AddCounterWaitGroup(int) adds the number of go routines to the wait group. This is useful when you want to wait for
	// all the go routines to finish before you get the final value of the counter.
	// This works as Add(int) method of sync.WaitGroup.
	AddCounterWaitGroup(int)
	// Done() decrements the wait group counter by 1. This is useful when you want to wait for
	// all the go routines to finish before you get the final value of the counter.
	// This works as Done() method of sync.WaitGroup.
	// Each go routine that calls AddCounterWaitGroup(int) should call Done() method,
	// when it no longer updates the counter value.
	Done()
	// GetFinalValue() returns the final value of the counter. This method waits for all the go routines
	// to finish before returning the final value of the counter.
	GetFinalValue() T
	// Freeze() freezes the counter value until Release() is not called. Any go routine that calls Increament(), Decreament(), IncrementBy(T) or DecrementBy(T)
	// will be blocked until Release() is called.
	Freeze()
	// Release() releases the counter value lock. After this method is called, any go routine can again update the counter value.
	Release()
}

func NewCounter[T incrementable](initialValue T) Counter[T] {
	return &counter[T]{
		count: initialValue,
		lck:   &sync.Mutex{},
		wg:    &sync.WaitGroup{},
	}
}

type counter[T incrementable] struct {
	count T
	lck   *sync.Mutex
	wg    *sync.WaitGroup
}

func (c *counter[T]) Increament() {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.count++
}

func (c *counter[T]) Decreament() {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.count--
}

func (c *counter[T]) GetCount() T {
	c.lck.Lock()
	defer c.lck.Unlock()
	return c.count
}

func (c *counter[T]) IncrementBy(n T) {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.count += n
}

func (c *counter[T]) DecrementBy(n T) {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.count -= n
}

func (c *counter[T]) Reset() {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.count = 0
}

func (c *counter[T]) SetCount(n T) {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.count = n
}

func (c *counter[T]) Freeze() {
	c.lck.Lock()
}

func (c *counter[T]) Release() {
	c.lck.Unlock()
}

func (c *counter[T]) GetCountAndReset() T {
	c.lck.Lock()
	defer c.lck.Unlock()
	count := c.count
	c.count = 0
	return count
}

func (c *counter[T]) AddCounterWaitGroup(i int) {
	c.wg.Add(i)
}

func (c *counter[T]) CounterWait() {
	c.wg.Wait()
}

func (c *counter[T]) GetFinalValue() T {
	c.wg.Wait()
	return c.GetCount()

}

func (c *counter[T]) Done() {
	c.wg.Done()
}
