package collection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayBinaryTree_Insert(t *testing.T) {
	binaryTree := NewArrayBinaryTree[int]()
	count := 15
	for i := 0; i < count; i++ {
		binaryTree.Insert(i)
	}
	require.Equal(t, count, binaryTree.Size())
}

func TestArrayBinaryTree_InorderTraversal(t *testing.T) {
	t.Run("traverse an empty tree", func(t *testing.T) {
		binaryTree := NewArrayBinaryTree[int]()
		require.Equal(t, []int{}, binaryTree.InorderTraversal())
	})

	t.Run("traverse a tree", func(t *testing.T) {
		binaryTree := NewArrayBinaryTree[int]()
		//         0
		//       /    \
		//      1      2
		//    /   \   / \
		//   3     4 5   6
		//  / \   /
		// 7   8 9
		for i := 0; i < 10; i++ {
			binaryTree.Insert(i)
		}

		require.Equal(t, 10, binaryTree.Size())
		require.Equal(t, []int{7, 3, 8, 1, 9, 4, 0, 5, 2, 6}, binaryTree.InorderTraversal())
	})
}

func TestArrayBinaryTree_LevelOrderTraversal(t *testing.T) {
	t.Run("traverse an empty tree", func(t *testing.T) {
		binaryTree := NewArrayBinaryTree[int]()
		require.Equal(t, []int{}, binaryTree.LevelOrderTraversal())
	})

	t.Run("traverse a tree", func(t *testing.T) {
		binaryTree := NewArrayBinaryTree[int]()
		//         0
		//       /    \
		//      1      2
		//    /   \   / \
		//   3     4 5   6
		//  / \   /
		// 7   8 9
		for i := 0; i < 10; i++ {
			binaryTree.Insert(i)
		}

		require.Equal(t, 10, binaryTree.Size())
		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, binaryTree.LevelOrderTraversal())
	})
}
func TestArrayBinaryTree_PreorderTraversal(t *testing.T) {
	t.Run("traverse an empty tree", func(t *testing.T) {
		binaryTree := NewArrayBinaryTree[int]()
		require.Equal(t, []int{}, binaryTree.PreorderTraversal())
	})

	t.Run("traverse a tree", func(t *testing.T) {
		binaryTree := NewArrayBinaryTree[int]()
		//         0
		//       /    \
		//      1      2
		//    /   \   / \
		//   3     4 5   6
		//  / \   /
		// 7   8 9
		for i := 0; i < 10; i++ {
			binaryTree.Insert(i)
		}

		require.Equal(t, 10, binaryTree.Size())
		require.Equal(t, []int{0, 1, 3, 7, 8, 4, 9, 2, 5, 6}, binaryTree.PreorderTraversal())
	})
}

func TestArrayBinaryTree_PostorderTraversal(t *testing.T) {
	t.Run("traverse an empty tree", func(t *testing.T) {
		binaryTree := NewArrayBinaryTree[int]()
		require.Equal(t, []int{}, binaryTree.PostorderTraversal())
	})

	t.Run("traverse a tree", func(t *testing.T) {
		binaryTree := NewArrayBinaryTree[int]()
		//         0
		//       /    \
		//      1      2
		//    /   \   / \
		//   3     4 5   6
		//  / \   /
		// 7   8 9
		for i := 0; i < 10; i++ {
			binaryTree.Insert(i)
		}

		require.Equal(t, 10, binaryTree.Size())
		require.Equal(t, []int{7, 8, 3, 9, 4, 1, 5, 6, 2, 0}, binaryTree.PostorderTraversal())
	})
}
