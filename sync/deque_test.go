package sync

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceDeque_Push(t *testing.T) {
	deque := NewConcurrentSliceDeque[int]()
	val := 42
	deque.Push(val)

	require.False(t, deque.Empty(), "sliceDeque should not be empty after a push")
	require.Equal(t, 1, deque.Size(), "sliceDeque size should be 1 after one push")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val, peekVal, "Peek should return the last pushed value")
}

func TestConcurrentQueue_Enqueue(t *testing.T) {
	t.Parallel()
	q := NewConcurrentSliceDeque[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	require.Equal(t, 2, q.Size())
}

func TestSliceDeque_Pop(t *testing.T) {
	t.Run("pop from an empty deque should be false", func(t *testing.T) {
		deque := NewConcurrentSliceDeque[int]()
		actual, ok := deque.Pop()
		require.False(t, ok, "Pop on an empty deque should return false")
		require.Equal(t, 0, actual, "Pop on an empty deque should return zero value for int")
	})

	t.Run("pop from an non-empty deque should be true", func(t *testing.T) {
		deque := NewConcurrentSliceDeque[int]()
		val1, val2 := 1, 2
		deque.Push(val1)
		deque.Push(val2)

		poppedVal, ok := deque.Pop()
		require.True(t, ok, "Pop should return true for non-empty deque")
		require.Equal(t, val2, poppedVal, "Pop should return the last pushed value")
		require.Equal(t, 1, deque.Size(), "sliceDeque size should be 1 after one pop")

		peekVal, ok := deque.Peek()
		require.True(t, ok, "Peek should return true for non-empty deque")
		require.Equal(t, val1, peekVal, "Peek should return the second-to-last value after one pop")
	})
}

func TestSliceDeque_Peek(t *testing.T) {
	t.Run("Peek from an empty deque should be false", func(t *testing.T) {
		deque := NewConcurrentSliceDeque[int]()
		actual, ok := deque.Peek()
		require.False(t, ok, "Peek on an empty deque should return false")
		require.Equal(t, 0, actual, "Peek on an empty deque should return zero value for int")
	})

	t.Run("Peek from an non-empty deque should be true", func(t *testing.T) {
		deque := NewConcurrentSliceDeque[int]()
		deque.Push(1)

		actual, ok := deque.Peek()
		require.True(t, ok, "Peek should return true for non-empty deque")
		require.Equal(t, 1, actual, "Peek should return the last pushed value")
	})
}

func TestConcurrentQueue_Dequeue(t *testing.T) {
	t.Parallel()
	q := NewConcurrentSliceDeque[int]()
	q.Enqueue(1)
	q.Enqueue(2)

	val, ok := q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 1, val)

	val, ok = q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 2, val)

	_, ok = q.Dequeue()
	require.False(t, ok)
}

func TestConcurrentQueue_Front(t *testing.T) {
	t.Parallel()
	q := NewConcurrentSliceDeque[int]()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	val, ok := q.Front()
	require.True(t, ok)
	require.Equal(t, 0, val)

	q.Dequeue()

	val, ok = q.Front()
	require.True(t, ok)
	require.Equal(t, 1, val)
}

func TestConcurrentQueue_Back(t *testing.T) {
	t.Parallel()
	q := NewConcurrentSliceDeque[int]()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	val, ok := q.Back()
	require.True(t, ok)
	require.Equal(t, 9, val)

	q.Dequeue()

	val, ok = q.Back()
	require.True(t, ok)
	require.Equal(t, 9, val)

	q.Enqueue(10)
	val, ok = q.Back()
	require.True(t, ok)
	require.Equal(t, 10, val)
}

func TestConcurrentQueue_Empty(t *testing.T) {
	t.Parallel()
	q := NewConcurrentSliceDeque[int]()
	require.True(t, q.Empty())

	q.Enqueue(1)
	require.False(t, q.Empty())

	q.Dequeue()
	require.True(t, q.Empty())
}

func TestConcurrentQueue_Clear(t *testing.T) {
	t.Parallel()
	q := NewConcurrentSliceDeque[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Clear()

	require.Equal(t, 0, q.Size())
	require.True(t, q.Empty())
}

func TestConcurrentQueue_Contains(t *testing.T) {
	t.Parallel()
	q := NewConcurrentSliceDeque[int]()
	q.Enqueue(1)
	q.Enqueue(2)

	require.True(t, q.Contains(1))
	require.True(t, q.Contains(2))
	require.False(t, q.Contains(3))
}

func TestConcurrentQueue_Reverse(t *testing.T) {
	t.Parallel()
	q := NewConcurrentSliceDeque[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Reverse()

	val, ok := q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 3, val)

	val, ok = q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 2, val)

	val, ok = q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 1, val)
}

func TestConcurrentQueue_ToSlice(t *testing.T) {
	t.Parallel()
	q := NewConcurrentSliceDeque[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	require.ElementsMatch(t, []int{1, 2, 3}, q.ToSlice())
}

func TestSliceDeque_Top(t *testing.T) {
	t.Parallel()
	deque := NewConcurrentSliceDeque[int]()
	deque.Push(1)

	topVal, ok := deque.Top()
	require.True(t, ok, "Top should return true for non-empty deque")
	require.Equal(t, 1, topVal, "Top should return the last pushed value")
}

func TestSliceDeque_Front(t *testing.T) {
	t.Parallel()
	t.Run("front from an empty deque should return false", func(t *testing.T) {
		deque := NewConcurrentSliceDeque[int]()
		frontVal, ok := deque.Front()
		require.False(t, ok, "Front on an empty deque")
		require.Equal(t, 0, frontVal, "Front on an empty deque")
	})

	t.Run("front from an non-empty deque should be true", func(t *testing.T) {
		deque := NewConcurrentSliceDeque[int]()
		deque.Enqueue(1)
		deque.Enqueue(2)

		frontVal, ok := deque.Front()
		require.True(t, ok, "Front should return true for non-empty deque")
		require.Equal(t, 1, frontVal, "Front should return the first enqueued value")
	})
}

func TestSliceDeque_Back(t *testing.T) {
	t.Parallel()
	t.Run("back from an empty deque should return false", func(t *testing.T) {
		deque := NewConcurrentSliceDeque[int]()
		backVal, ok := deque.Back()
		require.False(t, ok, "Back on an empty deque")
		require.Equal(t, 0, backVal, "Back on an empty deque")
	})

	t.Run("back from an non-empty deque should be true", func(t *testing.T) {
		deque := NewConcurrentSliceDeque[int]()
		deque.Enqueue(1)
		deque.Enqueue(2)

		backVal, ok := deque.Back()
		require.True(t, ok, "Back should return true for non-empty deque")
		require.Equal(t, 2, backVal, "Back should return the last enqueued value")
	})
}

func TestConcurrentQueue_Size(t *testing.T) {
	t.Parallel()
	q := NewConcurrentSliceDeque[int]()
	require.Equal(t, 0, q.Size())

	q.Enqueue(1)
	require.Equal(t, 1, q.Size())

	q.Dequeue()
	require.Equal(t, 0, q.Size())
}

func TestConcurrentQueue_Remove(t *testing.T) {
	t.Parallel()
	t.Run("Remove an existing element", func(t *testing.T) {
		q := NewConcurrentSliceDeque[int]()
		for i := 0; i < 3; i++ {
			q.Enqueue(i)
		}

		ok := q.Remove(1)
		require.True(t, ok)

		require.ElementsMatch(t, []int{0, 2}, q.ToSlice())
	})

	t.Run("Remove an non-existing element", func(t *testing.T) {
		q := NewConcurrentSliceDeque[int]()
		for i := 0; i < 3; i++ {
			q.Enqueue(i)
		}

		ok := q.Remove(4)
		require.False(t, ok)

		require.ElementsMatch(t, []int{0, 1, 2}, q.ToSlice())
	})
}
