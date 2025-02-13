package datastructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStackQueue(t *testing.T) {
	stack := NewStackQueue[int]()
	t.Run("New stack should be empty", func(t *testing.T) {
		require.True(t, stack.Empty(), "New stack should be empty")
	})

	t.Run("New stack should have size 0", func(t *testing.T) {
		require.Equal(t, 0, stack.Size(), "New stack should have size 0")
	})

	t.Run("Peek on empty stack should return false", func(t *testing.T) {
		peekVal, ok := stack.Peek()
		require.False(t, ok, "Peek on empty stack should return false")
		require.Equal(t, 0, peekVal, "Peek on empty stack should return zero value for int")
	})

	t.Run("Pop on empty stack should return false", func(t *testing.T) {
		popVal, ok := stack.Pop()
		require.False(t, ok, "Pop on empty stack should return false")
		require.Equal(t, 0, popVal, "Pop on empty stack should return zero value for int")
	})

	t.Run("Dequeue on empty stack should return false", func(t *testing.T) {
		dequeueVal, ok := stack.Dequeue()
		require.False(t, ok, "Dequeue on empty stack should return false")
		require.Equal(t, 0, dequeueVal, "Dequeue on empty stack should return zero value for int")
	})

	t.Run("Front on stack should return true", func(t *testing.T) {
		frontVal, ok := stack.Front()
		require.False(t, ok, "Front on empty stack should return false")
		require.Equal(t, 0, frontVal, "Front on empty stack should return zero value for int")
	})

	t.Run("Back on stack should return true", func(t *testing.T) {
		backVal, ok := stack.Back()
		require.False(t, ok, "Back on empty stack should return false")
		require.Equal(t, 0, backVal, "Back on empty stack should return zero value for int")
	})
}

func TestPushAddsElementToStackQueue(t *testing.T) {
	stack := NewStackQueue[int]()
	val := 42
	stack.Push(val)

	require.False(t, stack.Empty(), "StackQueue should not be empty after a push")
	require.Equal(t, 1, stack.Size(), "StackQueue size should be 1 after one push")

	peekVal, ok := stack.Peek()
	require.True(t, ok, "Peek should return true for non-empty stackQueue")
	require.Equal(t, val, peekVal, "Peek should return the last pushed value")
}

func TestEnqueueAddsElementToStackQueue(t *testing.T) {
	stack := NewStackQueue[int]()
	val := 42
	stack.Enqueue(val)

	require.False(t, stack.Empty(), "StackQueue should not be empty after an enqueue")
	require.Equal(t, 1, stack.Size(), "StackQueue size should be 1 after one enqueue")

	peekVal, ok := stack.Peek()
	require.True(t, ok, "Peek should return true for non-empty stackQueue")
	require.Equal(t, val, peekVal, "Peek should return the last enqueued value")
}

func TestPopRemovesElementFromStackQueue(t *testing.T) {
	stack := NewStackQueue[int]()
	val1, val2 := 1, 2
	stack.Push(val1)
	stack.Push(val2)

	poppedVal, ok := stack.Pop()
	require.True(t, ok, "Pop should return true for non-empty stackQueue")
	require.Equal(t, val2, poppedVal, "Pop should return the last pushed value")
	require.Equal(t, 1, stack.Size(), "StackQueue size should be 1 after one pop")

	peekVal, ok := stack.Peek()
	require.True(t, ok, "Peek should return true for non-empty stackQueue")
	require.Equal(t, val1, peekVal, "Peek should return the second-to-last value after one pop")
}

func TestDequeueRemovesElementFromStackQueue(t *testing.T) {
	stack := NewStackQueue[int]()
	val1, val2 := 1, 2
	stack.Enqueue(val1)
	stack.Enqueue(val2)

	dequeuedVal, ok := stack.Dequeue()
	require.True(t, ok, "Dequeue should return true for non-empty stackQueue")
	require.Equal(t, val1, dequeuedVal, "Dequeue should return the first enqueued value")
	require.Equal(t, 1, stack.Size(), "StackQueue size should be 1 after one dequeue")

	peekVal, ok := stack.Peek()
	require.True(t, ok, "Peek should return true for non-empty stackQueue")
	require.Equal(t, val2, peekVal, "Peek should return the second enqueued value after one dequeue")
}

func TestClearStackQueue(t *testing.T) {
	stack := NewStackQueue[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Clear()

	require.True(t, stack.Empty(), "StackQueue should be empty after Clear()")
	require.Equal(t, 0, stack.Size(), "StackQueue size should be 0 after Clear()")
}

func TestContainsElement(t *testing.T) {
	stack := NewStackQueue[int]()
	stack.Push(1)
	stack.Push(2)

	require.True(t, stack.Contains(1), "StackQueue should contain the element 1")
	require.False(t, stack.Contains(3), "StackQueue should not contain the element 3")
}

func TestIndexOfElement(t *testing.T) {
	stack := NewStackQueue[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	index, ok := stack.IndexOf(2)
	require.True(t, ok, "StackQueue should contain the element 2")
	require.Equal(t, 1, index, "IndexOf should return the correct index for element 2")

	index, ok = stack.IndexOf(4)
	require.False(t, ok, "StackQueue should not contain the element 4")
	require.Equal(t, -1, index, "IndexOf should return -1 for element 4")
}

func TestReverseStackQueue(t *testing.T) {
	stack := NewStackQueue[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Reverse()

	// Check if the order is reversed
	peekVal, ok := stack.Peek()
	require.True(t, ok, "Peek should return true for non-empty stackQueue")
	require.Equal(t, 1, peekVal, "Peek should return the first pushed value after Reverse()")
}

func TestToSlice(t *testing.T) {
	stack := NewStackQueue[int]()
	stack.Push(1)
	stack.Push(2)

	slice := stack.ToSlice()
	require.Equal(t, []int{1, 2}, slice, "ToSlice should return a copy of the data")
}

func TestTopAlias(t *testing.T) {
	stack := NewStackQueue[int]()
	stack.Push(1)

	topVal, ok := stack.Top()
	require.True(t, ok, "Top should return true for non-empty stackQueue")
	require.Equal(t, 1, topVal, "Top should return the last pushed value")
}

func TestFront(t *testing.T) {
	stack := NewStackQueue[int]()
	stack.Enqueue(1)
	stack.Enqueue(2)

	frontVal, ok := stack.Front()
	require.True(t, ok, "Front should return true for non-empty stackQueue")
	require.Equal(t, 1, frontVal, "Front should return the first enqueued value")
}

func TestBack(t *testing.T) {
	stack := NewStackQueue[int]()
	stack.Enqueue(1)
	stack.Enqueue(2)

	backVal, ok := stack.Back()
	require.True(t, ok, "Back should return true for non-empty stackQueue")
	require.Equal(t, 2, backVal, "Back should return the last enqueued value")
}

func TestRemoveAt(t *testing.T) {
	t.Run("Remove from a stackQueue", func(t *testing.T) {
		stack := NewStackQueue[int]()
		for i := 0; i < 3; i++ {
			stack.Push(i)
		}

		removedVal, ok := stack.RemoveAt(1)
		require.True(t, ok, "RemoveAt should return true for valid index")
		require.Equal(t, 1, removedVal, "RemoveAt should return the correct value")

		require.Equal(t, 2, stack.Size(), "StackQueue size should be 2 after RemoveAt")
		require.Equal(t, []int{0, 2}, stack.ToSlice(), "RemoveAt should remove the correct element")
	})

	t.Run("Remove with a negative index", func(t *testing.T) {
		stack := NewStackQueue[int]()
		for i := 0; i < 3; i++ {
			stack.Push(i)
		}
		removedVal, ok := stack.RemoveAt(-1)
		require.False(t, ok, "RemoveAt should return false for remove at")
		require.Equal(t, 0, removedVal, "RemoveAt should return zero value")
	})

	t.Run("Remove with an invalid index", func(t *testing.T) {
		stack := NewStackQueue[int]()
		for i := 0; i < 3; i++ {
			stack.Push(i)
		}

		removedVal, ok := stack.RemoveAt(3)
		require.False(t, ok, "RemoveAt should return false for remove at")
		require.Equal(t, 0, removedVal, "RemoveAt should return zero value")
	})
}
