package counter

import "sync"

type incrementable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

type Counter[T any] interface {
	AddCounterWaitGroup(int)
	CounterWait()
	Done()
	Increament()
	Decreament()
	GetCount() T
	GetCountAndReset() T
	Freeze()
	Release()
	IncrementBy(T)
	DecrementBy(T)
	Reset()
	SetCount(T)
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

func (c *counter[T]) Done() {
	c.wg.Done()
}
