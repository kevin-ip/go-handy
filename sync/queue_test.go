package sync

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcurrentQueue_Enqueue(t *testing.T) {
	t.Parallel()
	q := NewConcurrentQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	require.Equal(t, 2, q.Size())
}

func TestConcurrentQueue_Dequeue(t *testing.T) {
	t.Parallel()
	q := NewConcurrentQueue[int]()
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
	q := NewConcurrentQueue[int]()
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
	q := NewConcurrentQueue[int]()
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
	q := NewConcurrentQueue[int]()
	require.True(t, q.Empty())

	q.Enqueue(1)
	require.False(t, q.Empty())

	q.Dequeue()
	require.True(t, q.Empty())
}

func TestConcurrentQueue_Size(t *testing.T) {
	t.Parallel()
	q := NewConcurrentQueue[int]()
	require.Equal(t, 0, q.Size())

	q.Enqueue(1)
	require.Equal(t, 1, q.Size())

	q.Dequeue()
	require.Equal(t, 0, q.Size())
}

func TestConcurrentQueue_Clear(t *testing.T) {
	t.Parallel()
	q := NewConcurrentQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Clear()

	require.Equal(t, 0, q.Size())
	require.True(t, q.Empty())
}

func TestConcurrentQueue_Contains(t *testing.T) {
	t.Parallel()
	q := NewConcurrentQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)

	require.True(t, q.Contains(1))
	require.True(t, q.Contains(2))
	require.False(t, q.Contains(3))
}

func TestConcurrentQueue_Reverse(t *testing.T) {
	t.Parallel()
	q := NewConcurrentQueue[int]()
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
	q := NewConcurrentQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	require.ElementsMatch(t, []int{1, 2, 3}, q.ToSlice())
}

func TestConcurrentQueue_Remove(t *testing.T) {
	t.Parallel()
	t.Run("Remove an existing element", func(t *testing.T) {
		q := NewConcurrentQueue[int]()
		for i := 0; i < 3; i++ {
			q.Enqueue(i)
		}

		ok := q.Remove(1)
		require.True(t, ok)

		require.ElementsMatch(t, []int{0, 2}, q.ToSlice())
	})

	t.Run("Remove an non-existing element", func(t *testing.T) {
		q := NewConcurrentQueue[int]()
		for i := 0; i < 3; i++ {
			q.Enqueue(i)
		}

		ok := q.Remove(4)
		require.False(t, ok)

		require.ElementsMatch(t, []int{0, 1, 2}, q.ToSlice())
	})
}
