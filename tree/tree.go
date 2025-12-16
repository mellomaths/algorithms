package tree

type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}
