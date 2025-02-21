package sync

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kevin-ip/go-handy/collection"
	"github.com/kevin-ip/go-handy/collectiontest"
)

func TestConcurrentSliceDeque(t *testing.T) {
	newDequeFunc := func() collection.Deque[int] {
		return NewConcurrentSliceDeque[int]()
	}
	collectiontest.RunDequeTests(t, newDequeFunc)
	runConcurrentDequeTest(t, newDequeFunc)
}

func TestConcurrentLinkedDeque(t *testing.T) {
	newDequeFunc := func() collection.Deque[int] {
		return NewConcurrentLinkedDeque[int]()
	}
	collectiontest.RunDequeTests(t, newDequeFunc)
	runConcurrentDequeTest(t, newDequeFunc)
}

func runConcurrentDequeTest(t *testing.T, newDequeFunc func() collection.Deque[int]) {
	testCases := []struct {
		name string
		test func(t *testing.T, newDequeFunc func() collection.Deque[int])
	}{
		{"testConcurrentDeque_Front", testConcurrentDeque_Front},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.test(t, newDequeFunc)
		})
	}
}

func testConcurrentDeque_Front(t *testing.T, newDequeFunc func() collection.Deque[int]) {
	t.Parallel()
	q := newDequeFunc()
	for i := 0; i < 3; i++ {
		q.Enqueue(i)
	}

	for i, expected := range []bool{true, true, true, false} {
		val, ok := q.Front()
		require.Equal(t, expected, ok)
		if expected {
			require.Equal(t, i, val)
		}

		q.Dequeue()
	}
}
