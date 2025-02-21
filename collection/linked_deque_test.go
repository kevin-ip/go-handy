package collection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkedDeque_Push(t *testing.T) {
	deque := NewLinkedDeque[int]()
	val := 42
	deque.Push(val)

	require.False(t, deque.Empty(), "Deque should not be empty after a push")
	require.Equal(t, 1, deque.Size(), "Deque size should be 1 after one push")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val, peekVal, "Peek should return the last pushed value")
}

func TestLinkedDeque_Enqueue(t *testing.T) {
	deque := NewLinkedDeque[int]()
	val := 42
	deque.Enqueue(val)

	require.False(t, deque.Empty(), "Deque should not be empty after an enqueue")
	require.Equal(t, 1, deque.Size(), "Deque size should be 1 after one enqueue")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val, peekVal, "Peek should return the last enqueued value")
}

func TestLinkedDeque_Pop(t *testing.T) {
	t.Run("pop from an empty deque should be false", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		actual, ok := deque.Pop()
		require.False(t, ok, "Pop on an empty deque should return false")
		require.Equal(t, 0, actual, "Pop on an empty deque should return zero value for int")
	})

	t.Run("pop from an non-empty deque should be true", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
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
	})

	t.Run("pop all elements should be empty", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		val1, val2 := 1, 2
		deque.Push(val1)
		deque.Push(val2)

		poppedVal1, ok := deque.Pop()
		require.True(t, ok, "Pop should return true for non-empty deque")
		require.Equal(t, val2, poppedVal1, "Pop should return the last pushed value")

		poppedVal2, ok := deque.Pop()
		require.True(t, ok, "Pop should return true for non-empty deque")
		require.Equal(t, val1, poppedVal2, "Pop should return the last pushed value")

		poppedVal3, ok := deque.Pop()
		require.False(t, ok, "Pop on an empty deque should return false")
		require.Equal(t, 0, poppedVal3, "Pop on an empty deque should return zero value for int")
	})
}

func TestLinkedDeque_Peek(t *testing.T) {
	t.Run("Peek from an empty deque should be false", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		actual, ok := deque.Peek()
		require.False(t, ok, "Peek on an empty deque should return false")
		require.Equal(t, 0, actual, "Peek on an empty deque should return zero value for int")
	})

	t.Run("Peek from an non-empty deque should be true", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		deque.Push(1)

		actual, ok := deque.Peek()
		require.True(t, ok, "Peek should return true for non-empty deque")
		require.Equal(t, 1, actual, "Peek should return the last pushed value")
	})
}

func TestLinkedDeque_Dequeue(t *testing.T) {
	t.Run("dequeue from an empty deque should be false", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		actual, ok := deque.Dequeue()
		require.False(t, ok, "Dequeue on an empty deque should return false")
		require.Equal(t, 0, actual, "Dequeue on an empty deque")
	})

	t.Run("dequeue from an non-empty deque should be true", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
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
	})
}

func TestLinkedDeque_Clear(t *testing.T) {
	deque := NewLinkedDeque[int]()
	deque.Push(1)
	deque.Push(2)
	deque.Clear()

	require.True(t, deque.Empty(), "Deque should be empty after Clear()")
	require.Equal(t, 0, deque.Size(), "Deque size should be 0 after Clear()")
}

func TestLinkedDeque_Contains(t *testing.T) {
	deque := NewLinkedDeque[int]()
	deque.Push(1)
	deque.Push(2)

	require.True(t, deque.Contains(1), "Deque should contain the element 1")
	require.False(t, deque.Contains(3), "Deque should not contain the element 3")
}

func TestLinkedDeque_IndexOf(t *testing.T) {
	deque := NewLinkedDeque[int]()
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

func TestLinkedDeque_Reverse(t *testing.T) {
	t.Run("Reverse with odd number of elements", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}

		deque.Reverse()

		// Check if the order is reversed
		actual := deque.ToSlice()
		require.EqualValues(t, []int{2, 1, 0}, actual, "Deque should return reversed elements")
	})

	t.Run("Reverse with even number of elements", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		for i := 0; i < 4; i++ {
			deque.Push(i)
		}

		deque.Reverse()

		// Check if the order is reversed
		actual := deque.ToSlice()
		require.EqualValues(t, []int{3, 2, 1, 0}, actual, "Deque should return reversed elements")
	})

	t.Run("Reverse an empty deque", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		deque.Reverse()
		actual := deque.ToSlice()
		require.EqualValues(t, []int{}, actual, "Deque should return reversed elements")
	})
}

func TestLinkedDeque_ToSlice(t *testing.T) {
	deque := NewLinkedDeque[int]()
	deque.Push(1)
	deque.Push(2)

	slice := deque.ToSlice()
	require.Equal(t, []int{1, 2}, slice, "ToSlice should return a copy of the data")
}

func TestLinkedDeque_Top(t *testing.T) {
	deque := NewLinkedDeque[int]()
	deque.Push(1)

	topVal, ok := deque.Top()
	require.True(t, ok, "Top should return true for non-empty deque")
	require.Equal(t, 1, topVal, "Top should return the last pushed value")
}

func TestLinkedDeque_Front(t *testing.T) {
	t.Run("front from an empty deque should return false", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		frontVal, ok := deque.Front()
		require.False(t, ok, "Front on an empty deque")
		require.Equal(t, 0, frontVal, "Front on an empty deque")
	})

	t.Run("front from an non-empty deque should be true", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		deque.Enqueue(1)
		deque.Enqueue(2)

		frontVal, ok := deque.Front()
		require.True(t, ok, "Front should return true for non-empty deque")
		require.Equal(t, 1, frontVal, "Front should return the first enqueued value")
	})
}

func TestLinkedDeque_Back(t *testing.T) {
	t.Run("back from an empty deque should return false", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		backVal, ok := deque.Back()
		require.False(t, ok, "Back on an empty deque")
		require.Equal(t, 0, backVal, "Back on an empty deque")
	})

	t.Run("back from an non-empty deque should be true", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		deque.Enqueue(1)
		deque.Enqueue(2)

		backVal, ok := deque.Back()
		require.True(t, ok, "Back should return true for non-empty deque")
		require.Equal(t, 2, backVal, "Back should return the last enqueued value")
	})
}

func TestLinkedDeque_Size(t *testing.T) {
	t.Run("Size from an empty deque should be zero", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		require.Equal(t, 0, deque.Size(), "Deque size should be 0")
	})

	t.Run("Size from an non-empty deque should return the count of elements", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}
		require.Equal(t, 3, deque.Size(), "Deque size should be 3")
	})
}

func TestLinkedDeque_RemoveAt(t *testing.T) {
	t.Run("Remove from a Deque", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
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
		deque := NewLinkedDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}
		removedVal, ok := deque.RemoveAt(-1)
		require.False(t, ok, "RemoveAt should return false for remove at")
		require.Equal(t, 0, removedVal, "RemoveAt should return zero value")
	})

	t.Run("Remove with an invalid index", func(t *testing.T) {
		deque := NewLinkedDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}

		removedVal, ok := deque.RemoveAt(3)
		require.False(t, ok, "RemoveAt should return false for remove at")
		require.Equal(t, 0, removedVal, "RemoveAt should return zero value")
	})
}
