package collection

import (
	"cmp"
)

type ArrayBinaryTree[X cmp.Ordered] struct {
	tree []*X
	size int
}

func NewArrayBinaryTree[X cmp.Ordered]() BinaryTree[X] {
	return &ArrayBinaryTree[X]{
		tree: make([]*X, 10),
	}
}

func (a *ArrayBinaryTree[X]) Insert(x X) {
	deque := NewSliceDeque[int]()
	deque.Enqueue(0)

	for !deque.Empty() {
		index, ok := deque.Dequeue()
		if !ok {
			return
		}

		if index >= cap(a.tree) {
			a.expand()
		}

		currentNode := a.tree[index]
		if currentNode == nil {
			a.tree[index] = &x
			a.size++
			return
		} else {
			leftIndex := a.leftIndex(index)
			deque.Enqueue(leftIndex)
			rightIndex := a.rightIndex(index)
			deque.Enqueue(rightIndex)
		}
	}
}

func (a *ArrayBinaryTree[X]) leftIndex(index int) int {
	return index*2 + 1
}

func (a *ArrayBinaryTree[X]) rightIndex(index int) int {
	return index*2 + 2
}

func (a *ArrayBinaryTree[X]) parentIndex(index int) (int, bool) {
	if index == 0 {
		return -1, false
	}
	return (index - 1) / 2, true
}

func (a *ArrayBinaryTree[X]) Delete(x X) bool {
	// TODO implement me
	panic("implement me")
}

func (a *ArrayBinaryTree[X]) Size() int {
	return a.size
}

func (a *ArrayBinaryTree[X]) InorderTraversal() []X {
	return a.inorderTraversal(0)
}

func (a *ArrayBinaryTree[X]) inorderTraversal(index int) []X {
	if index >= cap(a.tree) || a.tree[index] == nil {
		return []X{}
	}

	left := a.inorderTraversal(a.leftIndex(index))
	right := a.inorderTraversal(a.rightIndex(index))
	result := []X{}
	result = append(result, left...)
	result = append(result, *a.tree[index])
	result = append(result, right...)
	return result
}

func (a *ArrayBinaryTree[X]) PreorderTraversal() []X {
	return a.preorderTraversal(0)
}

func (a *ArrayBinaryTree[X]) preorderTraversal(index int) []X {
	if index >= cap(a.tree) || a.tree[index] == nil {
		return []X{}
	}

	result := []X{*a.tree[index]}
	left := a.preorderTraversal(a.leftIndex(index))
	right := a.preorderTraversal(a.rightIndex(index))
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func (a *ArrayBinaryTree[X]) PostorderTraversal() []X {
	return a.postorderTraversal(0)
}

func (a *ArrayBinaryTree[X]) postorderTraversal(index int) []X {
	if index >= cap(a.tree) || a.tree[index] == nil {
		return []X{}
	}

	result := []X{}
	left := a.postorderTraversal(a.leftIndex(index))
	right := a.postorderTraversal(a.rightIndex(index))
	result = append(result, left...)
	result = append(result, right...)
	result = append(result, *a.tree[index])
	return result
}

func (a *ArrayBinaryTree[X]) expand() {
	newTree := make([]*X, cap(a.tree)*2)
	copy(newTree, a.tree)
	a.tree = newTree
}
