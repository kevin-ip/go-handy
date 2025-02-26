package collection

import "cmp"

type BinaryTree[X cmp.Ordered] interface {
	// Insert inserts the value into the binary tree at the first position available in level order manner.
	Insert(X)

	// Delete removes the value in the binary tree at the first position available in level order manner.
	Delete(X) bool

	// Size returns the total elements in the binary tree
	Size() int

	// InorderTraversal
	// Given the tree below, it returns 3, 1, 4, 0, 5, 2
	//         0
	//       /    \
	//      1      2
	//    /   \   /
	//   3     4 5
	InorderTraversal() []X

	// LevelOrderTraversal
	// Given the tree below, it returns 0, 1, 2, 3, 4, 5
	//         0
	//       /    \
	//      1      2
	//    /   \   /
	//   3     4 5
	LevelOrderTraversal() []X

	// PreorderTraversal
	// Given the tree below, it returns 0, 1, 3, 4, 2, 5
	//         0
	//       /    \
	//      1      2
	//    /   \   /
	//   3     4 5
	PreorderTraversal() []X

	// PostorderTraversal
	// Given the tree below, it returns 3, 4, 1, 5, 2, 0
	//         0
	//       /    \
	//      1      2
	//    /   \   /
	//   3     4 5
	PostorderTraversal() []X
}
