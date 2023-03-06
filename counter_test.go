package counter

import (
	"sync"
	"testing"
	"time"
)

func TestNewCounter(t *testing.T) {

	counter := NewCounter(2)
	if counter == nil {
		t.Errorf("Expected counter to be not nil")
	}

}

func TestIncreament(t *testing.T) {
	counter := NewCounter(2)
	counter.Increament()
	if counter.GetCount() != 3 {
		t.Errorf("Expected counter to be 3, but got %d", counter.GetCount())
	}
}

func TestDecreament(t *testing.T) {
	counter := NewCounter(2)
	counter.Decreament()
	if counter.GetCount() != 1 {
		t.Errorf("Expected counter to be 1, but got %d", counter.GetCount())
	}
}

func TestGetCount(t *testing.T) {
	counter := NewCounter(2)
	if counter.GetCount() != 2 {
		t.Errorf("Expected counter to be 2, but got %d", counter.GetCount())
	}
}

func TestIncrementBy(t *testing.T) {
	counter := NewCounter(5)
	counter.IncrementBy(3)
	if counter.GetCount() != 8 {
		t.Errorf("Expected counter to be 8, but got %d", counter.GetCount())
	}
}

func TestDecrementBy(t *testing.T) {
	counter := NewCounter(5)
	counter.DecrementBy(3)
	if counter.GetCount() != 2 {
		t.Errorf("Expected counter to be 2, but got %d", counter.GetCount())
	}
}

func TestReset(t *testing.T) {
	counter := NewCounter(5)
	counter.Reset()
	if counter.GetCount() != 0 {
		t.Errorf("Expected counter to be 0, but got %d", counter.GetCount())
	}
}

func TestSetCount(t *testing.T) {
	counter := NewCounter(6)
	counter.SetCount(9)
	if counter.GetCount() != 9 {
		t.Errorf("Expected counter to be 9, but got %d", counter.GetCount())
	}
}

func TestGetCountAndReset(t *testing.T) {
	counter := NewCounter(6)
	if counter.GetCountAndReset() != 6 {
		t.Errorf("Expected counter to be 6, but got %d", counter.GetCount())
	}

	if counter.GetCount() != 0 {
		t.Errorf("Expected counter to be 0, but got %d", counter.GetCount())
	}

}

func TestRaceCondition(t *testing.T) {

	counter := NewCounter(0)
	wg := sync.WaitGroup{}
	var result int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		counter.AddCounterWaitGroup(1)
		go func() {

			for j := 0; j < 1000; j++ {
				time.Sleep(time.Millisecond * 10)
				counter.Increament()
			}

			wg.Done()
			counter.Done()
		}()
	}
	timer := time.NewTicker(time.Millisecond * 3)

	go func() {
		for {

			select {
			case <-timer.C:
				result += counter.GetCountAndReset()
			case <-time.After(time.Second * 5):
				return
			}

		}
	}()

	wg.Wait()

	result += counter.GetFinalValue()

	if result != 10000 {
		t.Errorf("Expected counter to be 10000, but counter got %d result got %d", counter.GetCount(), result)
	}

}
