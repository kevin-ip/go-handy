package collection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceDeque_Push(t *testing.T) {
	deque := NewSliceDeque[int]()
	val := 42
	deque.Push(val)

	require.False(t, deque.Empty(), "sliceDeque should not be empty after a push")
	require.Equal(t, 1, deque.Size(), "sliceDeque size should be 1 after one push")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val, peekVal, "Peek should return the last pushed value")
}

func TestSliceDeque_Enqueue(t *testing.T) {
	deque := NewSliceDeque[int]()
	val := 42
	deque.Enqueue(val)

	require.False(t, deque.Empty(), "sliceDeque should not be empty after an enqueue")
	require.Equal(t, 1, deque.Size(), "sliceDeque size should be 1 after one enqueue")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val, peekVal, "Peek should return the last enqueued value")
}

func TestSliceDeque_Pop(t *testing.T) {
	t.Run("pop from an empty deque should be false", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		actual, ok := deque.Pop()
		require.False(t, ok, "Pop on an empty deque should return false")
		require.Equal(t, 0, actual, "Pop on an empty deque should return zero value for int")
	})

	t.Run("pop from an non-empty deque should be true", func(t *testing.T) {
		deque := NewSliceDeque[int]()
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
		deque := NewSliceDeque[int]()
		actual, ok := deque.Peek()
		require.False(t, ok, "Peek on an empty deque should return false")
		require.Equal(t, 0, actual, "Peek on an empty deque should return zero value for int")
	})

	t.Run("Peek from an non-empty deque should be true", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		deque.Push(1)

		actual, ok := deque.Peek()
		require.True(t, ok, "Peek should return true for non-empty deque")
		require.Equal(t, 1, actual, "Peek should return the last pushed value")
	})
}

func TestSliceDeque_Dequeue(t *testing.T) {
	t.Run("dequeue from an empty deque should be false", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		actual, ok := deque.Dequeue()
		require.False(t, ok, "Dequeue on an empty deque should return false")
		require.Equal(t, 0, actual, "Dequeue on an empty deque")
	})

	t.Run("dequeue from an non-empty deque should be true", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		val1, val2 := 1, 2
		deque.Enqueue(val1)
		deque.Enqueue(val2)

		dequeuedVal, ok := deque.Dequeue()
		require.True(t, ok, "Dequeue should return true for non-empty deque")
		require.Equal(t, val1, dequeuedVal, "Dequeue should return the first enqueued value")
		require.Equal(t, 1, deque.Size(), "sliceDeque size should be 1 after one dequeue")

		peekVal, ok := deque.Peek()
		require.True(t, ok, "Peek should return true for non-empty deque")
		require.Equal(t, val2, peekVal, "Peek should return the second enqueued value after one dequeue")
	})
}

func TestSliceDeque_Clear(t *testing.T) {
	deque := NewSliceDeque[int]()
	deque.Push(1)
	deque.Push(2)
	deque.Clear()

	require.True(t, deque.Empty(), "sliceDeque should be empty after Clear()")
	require.Equal(t, 0, deque.Size(), "sliceDeque size should be 0 after Clear()")
}

func TestSliceDeque_Contains(t *testing.T) {
	deque := NewSliceDeque[int]()
	deque.Push(1)
	deque.Push(2)

	require.True(t, deque.Contains(1), "sliceDeque should contain the element 1")
	require.False(t, deque.Contains(3), "sliceDeque should not contain the element 3")
}

func TestSliceDeque_Reverse(t *testing.T) {
	t.Run("Reverse with odd number of elements", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}

		deque.Reverse()

		// Check if the order is reversed
		actual := deque.ToSlice()
		require.EqualValues(t, []int{2, 1, 0}, actual, "sliceDeque should return reversed elements")
	})

	t.Run("Reverse with even number of elements", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		for i := 0; i < 4; i++ {
			deque.Push(i)
		}

		deque.Reverse()

		// Check if the order is reversed
		actual := deque.ToSlice()
		require.EqualValues(t, []int{3, 2, 1, 0}, actual, "sliceDeque should return reversed elements")
	})
}

func TestSliceDeque_ToSlice(t *testing.T) {
	deque := NewSliceDeque[int]()
	deque.Push(1)
	deque.Push(2)

	slice := deque.ToSlice()
	require.Equal(t, []int{1, 2}, slice, "ToSlice should return a copy of the data")
}

func TestSliceDeque_Top(t *testing.T) {
	deque := NewSliceDeque[int]()
	deque.Push(1)

	topVal, ok := deque.Top()
	require.True(t, ok, "Top should return true for non-empty deque")
	require.Equal(t, 1, topVal, "Top should return the last pushed value")
}

func TestSliceDeque_Front(t *testing.T) {
	t.Run("front from an empty deque should return false", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		frontVal, ok := deque.Front()
		require.False(t, ok, "Front on an empty deque")
		require.Equal(t, 0, frontVal, "Front on an empty deque")
	})

	t.Run("front from an non-empty deque should be true", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		deque.Enqueue(1)
		deque.Enqueue(2)

		frontVal, ok := deque.Front()
		require.True(t, ok, "Front should return true for non-empty deque")
		require.Equal(t, 1, frontVal, "Front should return the first enqueued value")
	})
}

func TestSliceDeque_Back(t *testing.T) {
	t.Run("back from an empty deque should return false", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		backVal, ok := deque.Back()
		require.False(t, ok, "Back on an empty deque")
		require.Equal(t, 0, backVal, "Back on an empty deque")
	})

	t.Run("back from an non-empty deque should be true", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		deque.Enqueue(1)
		deque.Enqueue(2)

		backVal, ok := deque.Back()
		require.True(t, ok, "Back should return true for non-empty deque")
		require.Equal(t, 2, backVal, "Back should return the last enqueued value")
	})
}

func TestSliceDeque_Size(t *testing.T) {
	t.Run("Size from an empty deque should be zero", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		require.Equal(t, 0, deque.Size(), "sliceDeque size should be 0")
	})

	t.Run("Size from an non-empty deque should return the count of elements", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}
		require.Equal(t, 3, deque.Size(), "sliceDeque size should be 3")
	})
}

func TestSliceDeque_RemoveAt(t *testing.T) {
	t.Run("Remove an existing value should be a success", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}

		ok := deque.Remove(1)
		require.True(t, ok, "Remove should return true")

		require.Equal(t, 2, deque.Size(), "sliceDeque size should be 2 after RemoveAt")
		require.Equal(t, []int{0, 2}, deque.ToSlice(), "RemoveAt should remove the correct element")
	})

	t.Run("Remove an existing value should fail", func(t *testing.T) {
		deque := NewSliceDeque[int]()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}
		ok := deque.Remove(-1)
		require.False(t, ok, "Remove should return false")
	})
}
