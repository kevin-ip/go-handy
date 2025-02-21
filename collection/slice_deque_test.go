package collection

import "testing"

func TestSliceDeque(t *testing.T) {
	runDequeTests(t, func() Deque[int] {
		return NewSliceDeque[int]()
	})
}
