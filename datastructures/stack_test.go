package datastructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	require.True(t, stack.Empty(), "New stack should be empty")
	require.Equal(t, 0, stack.Size(), "New stack should have size 0")

	val, ok := stack.Peek()
	require.False(t, ok, "Peek on empty stack should return false")
	require.Equal(t, 0, val, "Peek on empty stack should return zero value for int")
}

func TestPushAddsElementToStack(t *testing.T) {
	stack := NewStack[int]()
	val := 42
	stack.Push(val)

	require.False(t, stack.Empty(), "Stack should not be empty after a push")
	require.Equal(t, 1, stack.Size(), "Stack size should be 1 after one push")

	peekVal, ok := stack.Peek()
	require.True(t, ok, "Peek should return true for non-empty stack")
	require.Equal(t, val, peekVal, "Peek should return the last pushed value")
}

func TestPushMultipleElements(t *testing.T) {
	stack := NewStack[int]()
	val1, val2, val3 := 1, 2, 3
	stack.Push(val1)
	stack.Push(val2)
	stack.Push(val3)

	require.Equal(t, 3, stack.Size(), "Stack size should be 3 after three pushes")

	peekVal, ok := stack.Peek()
	require.True(t, ok, "Peek should return true for non-empty stack")
	require.Equal(t, val3, peekVal, "Peek should return the last pushed value")
}

func TestPopRemovesElementFromStack(t *testing.T) {
	stack := NewStack[int]()
	val1, val2 := 1, 2
	stack.Push(val1)
	stack.Push(val2)

	poppedVal, ok := stack.Pop()
	require.True(t, ok, "Pop should return true for non-empty stack")
	require.Equal(t, val2, poppedVal, "Pop should return the last pushed value")
	require.Equal(t, 1, stack.Size(), "Stack size should be 1 after one pop")

	peekVal, ok := stack.Peek()
	require.True(t, ok, "Peek should return true for non-empty stack")
	require.Equal(t, val1, peekVal, "Peek should return the second-to-last value after one pop")
}

func TestPopAllElements(t *testing.T) {
	stack := NewStack[int]()
	val1, val2 := 1, 2
	stack.Push(val1)
	stack.Push(val2)

	_, _ = stack.Pop() // Pop val2
	_, _ = stack.Pop() // Pop val1

	require.True(t, stack.Empty(), "Stack should be empty after all elements are popped")
	require.Equal(t, 0, stack.Size(), "Stack size should be 0 after all elements are popped")

	peekVal, ok := stack.Peek()
	require.False(t, ok, "Peek on empty stack should return false")
	require.Equal(t, 0, peekVal, "Peek on empty stack should return zero value for int")

	poppedVal, ok := stack.Pop()
	require.False(t, ok, "Pop on empty stack should return false")
	require.Equal(t, 0, poppedVal, "Pop on empty stack should return zero value for int")
}
