package design_practice

import "testing"

func TestMinStack(t *testing.T) {
	s := NewMinStack()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if s.GetMin() != -3 {
		t.Errorf("GetMin() = %d, want -3", s.GetMin())
	}
	s.Pop()
	if s.Top() != 0 {
		t.Errorf("Top() = %d, want 0", s.Top())
	}
	if s.GetMin() != -2 {
		t.Errorf("GetMin() after pop = %d, want -2", s.GetMin())
	}
}
