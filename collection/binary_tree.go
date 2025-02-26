package collection

import "cmp"

type BinaryTree[X cmp.Ordered] interface {
	// Insert inserts the value into the binary tree at the first position available in level order manner.
	Insert(X)

	// Delete removes the value in the binary tree at the first position available in level order manner.
	Delete(X) bool

	// Size returns the total elements in the binary tree
	Size() int

	InorderTraversal() []X
	PreorderTraversal() []X
	PostorderTraversal() []X
}
