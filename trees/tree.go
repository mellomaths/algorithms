package trees

import "errors"

var ErrEmptyTree = errors.New("empty tree, not able to calculate height")

type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}
