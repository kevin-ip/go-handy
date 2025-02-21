package collection

import (
	"testing"
)

func TestLinkedDeque(t *testing.T) {
	runDequeTests(t, func() Deque[int] {
		return NewLinkedDeque[int]()
	})
}
