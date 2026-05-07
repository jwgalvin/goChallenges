package min_stack

// MinStack is a stack that, in addition to the usual push/pop/top operations,
// can tell you the current minimum value in O(1) — no scanning the whole stack.
//
// The trick: every time you push a value, also remember what the minimum was
// at that moment. When you pop, that saved minimum comes back automatically.
// Think of it as each slot in the stack carrying two pieces of information:
// the value you pushed, and "what was the smallest thing in the stack right then."
type MinStack struct {
	vals []int
	mins []int
}

func NewMinStack() *MinStack {
	return &MinStack{}
}

func (s *MinStack) Push(val int) {
	s.vals = append(s.vals, val)
	if len(s.mins) == 0 || val < s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, val)
	} else {
		s.mins = append(s.mins, s.mins[len(s.mins)-1])
	}
}

func (s *MinStack) Pop() {
	s.vals = s.vals[:len(s.vals)-1]
	s.mins = s.mins[:len(s.mins)-1]
}

func (s *MinStack) Top() int {
	return s.vals[len(s.vals)-1]
}

func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}
