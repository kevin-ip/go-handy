package collection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func runDequeTests(t *testing.T, newDequeFunc func() Deque[int]) {
	testCases := []struct {
		name string
		test func(t *testing.T, newDequeFunc func() Deque[int])
	}{
		{"testDeque_Push", testDeque_Push},
		{"testDeque_Enqueue", testDeque_Enqueue},
		{"testDeque_Pop", testDeque_Pop},
		{"testDeque_Peek", testDeque_Peek},
		{"testDeque_Dequeue", testDeque_Dequeue},
		{"testDeque_Clear", testDeque_Clear},
		{"testDeque_Contains", testDeque_Contains},
		{"testDeque_Reverse", testDeque_Reverse},
		{"testDeque_ToSlice", testDeque_ToSlice},
		{"testDeque_Top", testDeque_Top},
		{"testDeque_Front", testDeque_Front},
		{"testDeque_Back", testDeque_Back},
		{"testDeque_Size", testDeque_Size},
		{"testDeque_Remove", testDeque_Remove},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.test(t, newDequeFunc)
		})
	}
}

func testDeque_Push(t *testing.T, newDequeFunc func() Deque[int]) {
	deque := newDequeFunc()
	val := 42
	deque.Push(val)

	require.False(t, deque.Empty(), "deque should not be empty after a push")
	require.Equal(t, 1, deque.Size(), "deque size should be 1 after one push")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val, peekVal, "Peek should return the last pushed value")
}

func testDeque_Enqueue(t *testing.T, newDequeFunc func() Deque[int]) {
	deque := newDequeFunc()
	val := 42
	deque.Enqueue(val)

	require.False(t, deque.Empty(), "deque should not be empty after an enqueue")
	require.Equal(t, 1, deque.Size(), "deque size should be 1 after one enqueue")

	peekVal, ok := deque.Peek()
	require.True(t, ok, "Peek should return true for non-empty deque")
	require.Equal(t, val, peekVal, "Peek should return the last enqueued value")
}

func testDeque_Pop(t *testing.T, newDequeFunc func() Deque[int]) {
	t.Run("pop from an empty deque should be false", func(t *testing.T) {
		deque := newDequeFunc()
		actual, ok := deque.Pop()
		require.False(t, ok, "Pop on an empty deque should return false")
		require.Equal(t, 0, actual, "Pop on an empty deque should return zero value for int")
	})

	t.Run("pop from an non-empty deque should be true", func(t *testing.T) {
		deque := newDequeFunc()
		val1, val2 := 1, 2
		deque.Push(val1)
		deque.Push(val2)

		poppedVal, ok := deque.Pop()
		require.True(t, ok, "Pop should return true for non-empty deque")
		require.Equal(t, val2, poppedVal, "Pop should return the last pushed value")
		require.Equal(t, 1, deque.Size(), "deque size should be 1 after one pop")
		require.Equal(t, []int{1}, deque.ToSlice())

		peekVal, ok := deque.Peek()
		require.True(t, ok, "Peek should return true for non-empty deque")
		require.Equal(t, val1, peekVal, "Peek should return the second-to-last value after one pop")

		poppedVal2, ok := deque.Pop()
		require.True(t, ok, "Pop should return true for non-empty deque")
		require.Equal(t, val1, poppedVal2, "Pop should return the last pushed value")
		require.Equal(t, 0, deque.Size(), "deque size should be 1 after one pop")
		require.Equal(t, []int{}, deque.ToSlice())

		_, ok = deque.Peek()
		require.False(t, ok, "Peek should return false for an empty deque")

		_, ok = deque.Pop()
		require.False(t, ok, "Pop should return false for an empty deque")

		_, ok = deque.Front()
		require.False(t, ok, "Front should return false for an empty deque")

		_, ok = deque.Back()
		require.False(t, ok, "Back should return false for an empty deque")
	})
}

func testDeque_Peek(t *testing.T, newDequeFunc func() Deque[int]) {
	t.Run("Peek from an empty deque should be false", func(t *testing.T) {
		deque := newDequeFunc()
		actual, ok := deque.Peek()
		require.False(t, ok, "Peek on an empty deque should return false")
		require.Equal(t, 0, actual, "Peek on an empty deque should return zero value for int")
	})

	t.Run("Peek from an non-empty deque should be true", func(t *testing.T) {
		deque := newDequeFunc()
		deque.Push(1)

		actual, ok := deque.Peek()
		require.True(t, ok, "Peek should return true for non-empty deque")
		require.Equal(t, 1, actual, "Peek should return the last pushed value")
	})
}

func testDeque_Dequeue(t *testing.T, newDequeFunc func() Deque[int]) {
	t.Run("dequeue from an empty deque should be false", func(t *testing.T) {
		deque := newDequeFunc()
		actual, ok := deque.Dequeue()
		require.False(t, ok, "Dequeue on an empty deque should return false")
		require.Equal(t, 0, actual, "Dequeue on an empty deque")
	})

	t.Run("dequeue from an non-empty deque should be true", func(t *testing.T) {
		deque := newDequeFunc()
		val1, val2 := 1, 2
		deque.Enqueue(val1)
		deque.Enqueue(val2)

		dequeuedVal, ok := deque.Dequeue()
		require.True(t, ok, "Dequeue should return true for non-empty deque")
		require.Equal(t, val1, dequeuedVal, "Dequeue should return the first enqueued value")
		require.Equal(t, 1, deque.Size(), "deque size should be 1 after one dequeue")
		require.Equal(t, []int{2}, deque.ToSlice())

		peekVal, ok := deque.Peek()
		require.True(t, ok, "Peek should return true for non-empty deque")
		require.Equal(t, val2, peekVal, "Peek should return the second enqueued value after one dequeue")

		dequeuedVal2, ok := deque.Dequeue()
		require.True(t, ok, "Dequeue should return true for non-empty deque")
		require.Equal(t, val2, dequeuedVal2, "Dequeue should return the first enqueued value")
		require.Equal(t, []int{}, deque.ToSlice())

		_, ok = deque.Peek()
		require.False(t, ok, "Peek should return false for an empty deque")

		_, ok = deque.Dequeue()
		require.False(t, ok, "Dequeue should return false for an empty deque")

		_, ok = deque.Front()
		require.False(t, ok, "Front should return false for an empty deque")

		_, ok = deque.Back()
		require.False(t, ok, "Back should return false for an empty deque")
	})
}

func testDeque_Clear(t *testing.T, newDequeFunc func() Deque[int]) {
	deque := newDequeFunc()
	deque.Push(1)
	deque.Push(2)
	deque.Clear()

	require.True(t, deque.Empty(), "deque should be empty after Clear()")
	require.Equal(t, 0, deque.Size(), "deque size should be 0 after Clear()")
}

func testDeque_Contains(t *testing.T, newDequeFunc func() Deque[int]) {
	deque := newDequeFunc()
	deque.Push(1)
	deque.Push(2)

	require.True(t, deque.Contains(1), "deque should contain the element 1")
	require.False(t, deque.Contains(3), "deque should not contain the element 3")
}

func testDeque_Reverse(t *testing.T, newDequeFunc func() Deque[int]) {
	t.Run("Reverse with odd number of elements", func(t *testing.T) {
		deque := newDequeFunc()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}

		deque.Reverse()

		// Check if the order is reversed
		actual := deque.ToSlice()
		require.EqualValues(t, []int{2, 1, 0}, actual, "deque should return reversed elements")
	})

	t.Run("Reverse with even number of elements", func(t *testing.T) {
		deque := newDequeFunc()
		for i := 0; i < 4; i++ {
			deque.Push(i)
		}

		deque.Reverse()

		// Check if the order is reversed
		actual := deque.ToSlice()
		require.EqualValues(t, []int{3, 2, 1, 0}, actual, "deque should return reversed elements")
	})

	t.Run("Reverse with no element", func(t *testing.T) {
		deque := newDequeFunc()
		deque.Reverse()
		actual := deque.ToSlice()
		require.EqualValues(t, []int{}, actual, "deque should return reversed elements")
	})
}

func testDeque_ToSlice(t *testing.T, newDequeFunc func() Deque[int]) {
	deque := newDequeFunc()
	deque.Push(1)
	deque.Push(2)

	slice := deque.ToSlice()
	require.Equal(t, []int{1, 2}, slice, "ToSlice should return a copy of the data")
}

func testDeque_Top(t *testing.T, newDequeFunc func() Deque[int]) {
	deque := newDequeFunc()
	deque.Push(1)

	topVal, ok := deque.Top()
	require.True(t, ok, "Top should return true for non-empty deque")
	require.Equal(t, 1, topVal, "Top should return the last pushed value")
}

func testDeque_Front(t *testing.T, newDequeFunc func() Deque[int]) {
	t.Run("front from an empty deque should return false", func(t *testing.T) {
		deque := newDequeFunc()
		frontVal, ok := deque.Front()
		require.False(t, ok, "Front on an empty deque")
		require.Equal(t, 0, frontVal, "Front on an empty deque")
	})

	t.Run("front from an non-empty deque should be true", func(t *testing.T) {
		deque := newDequeFunc()
		deque.Enqueue(1)
		deque.Enqueue(2)

		frontVal, ok := deque.Front()
		require.True(t, ok, "Front should return true for non-empty deque")
		require.Equal(t, 1, frontVal, "Front should return the first enqueued value")
	})
}

func testDeque_Back(t *testing.T, newDequeFunc func() Deque[int]) {
	t.Run("back from an empty deque should return false", func(t *testing.T) {
		deque := newDequeFunc()
		backVal, ok := deque.Back()
		require.False(t, ok, "Back on an empty deque")
		require.Equal(t, 0, backVal, "Back on an empty deque")
	})

	t.Run("back from an non-empty deque should be true", func(t *testing.T) {
		deque := newDequeFunc()
		deque.Enqueue(1)
		deque.Enqueue(2)

		backVal, ok := deque.Back()
		require.True(t, ok, "Back should return true for non-empty deque")
		require.Equal(t, 2, backVal, "Back should return the last enqueued value")
	})
}

func testDeque_Size(t *testing.T, newDequeFunc func() Deque[int]) {
	t.Run("Size from an empty deque should be zero", func(t *testing.T) {
		deque := newDequeFunc()
		require.Equal(t, 0, deque.Size(), "deque size should be 0")
	})

	t.Run("Size from an non-empty deque should return the count of elements", func(t *testing.T) {
		deque := newDequeFunc()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}
		require.Equal(t, 3, deque.Size(), "deque size should be 3")
	})
}

func testDeque_Remove(t *testing.T, newDequeFunc func() Deque[int]) {
	t.Run("Remove an existing value should be a success", func(t *testing.T) {
		deque := newDequeFunc()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}

		ok := deque.Remove(1)
		require.True(t, ok, "Remove should return true")

		require.Equal(t, 2, deque.Size(), "deque size should be 2 after RemoveAt")
		require.Equal(t, []int{0, 2}, deque.ToSlice())
		deque.Reverse()
		require.Equal(t, []int{2, 0}, deque.ToSlice())

		ok = deque.Remove(0)
		require.True(t, ok, "Remove should return true")
		require.Equal(t, []int{2}, deque.ToSlice())
		deque.Reverse()
		require.Equal(t, []int{2}, deque.ToSlice())

		ok = deque.Remove(2)
		require.True(t, ok, "Remove should return true")
		require.Equal(t, []int{}, deque.ToSlice())
		deque.Reverse()
		require.Equal(t, []int{}, deque.ToSlice())

		ok = deque.Remove(4)
		require.False(t, ok, "Remove should return false for an empty deque")

		_, ok = deque.Front()
		require.False(t, ok, "Front should return false for an empty deque")

		_, ok = deque.Back()
		require.False(t, ok, "Back should return false for an empty deque")
	})

	t.Run("Remove an existing value should fail", func(t *testing.T) {
		deque := newDequeFunc()
		for i := 0; i < 3; i++ {
			deque.Push(i)
		}
		ok := deque.Remove(-1)
		require.False(t, ok, "Remove should return false")
	})
}
