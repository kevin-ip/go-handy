package collection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDeque(t *testing.T) {
	deque := NewDeque[int]()
	t.Run("New deque should be empty", func(t *testing.T) {
		require.True(t, deque.Empty(), "New deque should be empty")
	})

	t.Run("New deque should have size 0", func(t *testing.T) {
		require.Equal(t, 0, deque.Size(), "New deque should have size 0")
	})

	t.Run("Peek on empty deque should return false", func(t *testing.T) {
		peekVal, ok := deque.Peek()
		require.False(t, ok, "Peek on empty deque should return false")
		require.Equal(t, 0, peekVal, "Peek on empty deque should return zero value for int")
	})

	t.Run("Pop on empty deque should return false", func(t *testing.T) {
		popVal, ok := deque.Pop()
		require.False(t, ok, "Pop on empty deque should return false")
		require.Equal(t, 0, popVal, "Pop on empty deque should return zero value for int")
	})

	t.Run("Dequeue on empty deque should return false", func(t *testing.T) {
		dequeueVal, ok := deque.Dequeue()
		require.False(t, ok, "Dequeue on empty deque should return false")
		require.Equal(t, 0, dequeueVal, "Dequeue on empty deque should return zero value for int")
	})

	t.Run("Front on deque should return true", func(t *testing.T) {
		frontVal, ok := deque.Front()
		require.False(t, ok, "Front on empty deque should return false")
		require.Equal(t, 0, frontVal, "Front on empty deque should return zero value for int")
	})

	t.Run("Back on deque should return true", func(t *testing.T) {
		backVal, ok := deque.Back()
		require.False(t, ok, "Back on empty deque should return false")
		require.Equal(t, 0, backVal, "Back on empty deque should return zero value for int")
	})
}

func TestPushAddsElementToDeque(t *testing.T) {
	deque := NewDeque[int]()
	val := 42
	deque.Push(val)

	require.False(t, deque.Empty(), "Deque should not be empty after a push")
	require.Equal(t, 1, deque.Size(), "Deque size should be 1 after one push")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val, peekVal, "Peek should return the last pushed value")
}

func TestEnqueueAddsElementToDeque(t *testing.T) {
	deque := NewDeque[int]()
	val := 42
	deque.Enqueue(val)

	require.False(t, deque.Empty(), "Deque should not be empty after an enqueue")
	require.Equal(t, 1, deque.Size(), "Deque size should be 1 after one enqueue")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val, peekVal, "Peek should return the last enqueued value")
}

func TestPopRemovesElementFromDeque(t *testing.T) {
	deque := NewDeque[int]()
	val1, val2 := 1, 2
	deque.Push(val1)
	deque.Push(val2)

	poppedVal, ok := deque.Pop()
	require.True(t, ok, "Pop should return true for non-empty deque")
	require.Equal(t, val2, poppedVal, "Pop should return the last pushed value")
	require.Equal(t, 1, deque.Size(), "Deque size should be 1 after one pop")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val1, peekVal, "Peek should return the second-to-last value after one pop")
}

func TestDequeueRemovesElementFromDeque(t *testing.T) {
	deque := NewDeque[int]()
	val1, val2 := 1, 2
	deque.Enqueue(val1)
	deque.Enqueue(val2)

	dequeuedVal, ok := deque.Dequeue()
	require.True(t, ok, "Dequeue should return true for non-empty deque")
	require.Equal(t, val1, dequeuedVal, "Dequeue should return the first enqueued value")
	require.Equal(t, 1, deque.Size(), "Deque size should be 1 after one dequeue")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val2, peekVal, "Peek should return the second enqueued value after one dequeue")
}

func TestClearDeque(t *testing.T) {
	deque := NewDeque[int]()
	deque.Push(1)
	deque.Push(2)
	deque.Clear()

	require.True(t, deque.Empty(), "Deque should be empty after Clear()")
	require.Equal(t, 0, deque.Size(), "Deque size should be 0 after Clear()")
}

func TestContainsElement(t *testing.T) {
	deque := NewDeque[int]()
	deque.Push(1)
	deque.Push(2)

	require.True(t, deque.Contains(1), "Deque should contain the element 1")
	require.False(t, deque.Contains(3), "Deque should not contain the element 3")
}

func TestIndexOfElement(t *testing.T) {
	deque := NewDeque[int]()
	deque.Push(1)
	deque.Push(2)
	deque.Push(3)

	index, ok := deque.IndexOf(2)
	require.True(t, ok, "Deque should contain the element 2")
	require.Equal(t, 1, index, "IndexOf should return the correct index for element 2")

	index, ok = deque.IndexOf(4)
	require.False(t, ok, "Deque should not contain the element 4")
	require.Equal(t, -1, index, "IndexOf should return -1 for element 4")
}

func TestReverseDeque(t *testing.T) {
	deque := NewDeque[int]()
	deque.Push(1)
	deque.Push(2)
	deque.Push(3)

	deque.Reverse()

	// Check if the order is reversed
	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, 1, peekVal, "Peek should return the first pushed value after Reverse()")
}

func TestToSlice(t *testing.T) {
	deque := NewDeque[int]()
	deque.Push(1)
	deque.Push(2)

	slice := deque.ToSlice()
	require.Equal(t, []int{1, 2}, slice, "ToSlice should return a copy of the data")
}

func TestTopAlias(t *testing.T) {
	deque := NewDeque[int]()
	deque.Push(1)

	topVal, ok := deque.Top()
	require.True(t, ok, "Top should return true for non-empty deque")
	require.Equal(t, 1, topVal, "Top should return the last pushed value")
}

func TestFront(t *testing.T) {
	deque := NewDeque[int]()
	deque.Enqueue(1)
	deque.Enqueue(2)

	frontVal, ok := deque.Front()
	require.True(t, ok, "Front should return true for non-empty deque")
	require.Equal(t, 1, frontVal, "Front should return the first enqueued value")
}

func TestBack(t *testing.T) {
	deque := NewDeque[int]()
	deque.Enqueue(1)
	deque.Enqueue(2)

	backVal, ok := deque.Back()
	require.True(t, ok, "Back should return true for non-empty deque")
	require.Equal(t, 2, backVal, "Back should return the last enqueued value")
}

func TestRemoveAt(t *testing.T) {
	t.Run("Remove from a Deque", func(t *testing.T) {
		deque := NewDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}

		removedVal, ok := deque.RemoveAt(1)
		require.True(t, ok, "RemoveAt should return true for valid index")
		require.Equal(t, 1, removedVal, "RemoveAt should return the correct value")

		require.Equal(t, 2, deque.Size(), "Deque size should be 2 after RemoveAt")
		require.Equal(t, []int{0, 2}, deque.ToSlice(), "RemoveAt should remove the correct element")
	})

	t.Run("Remove with a negative index", func(t *testing.T) {
		deque := NewDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}
		removedVal, ok := deque.RemoveAt(-1)
		require.False(t, ok, "RemoveAt should return false for remove at")
		require.Equal(t, 0, removedVal, "RemoveAt should return zero value")
	})

	t.Run("Remove with an invalid index", func(t *testing.T) {
		deque := NewDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}

		removedVal, ok := deque.RemoveAt(3)
		require.False(t, ok, "RemoveAt should return false for remove at")
		require.Equal(t, 0, removedVal, "RemoveAt should return zero value")
	})
}
