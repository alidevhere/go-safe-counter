package counter

import (
	"testing"
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
