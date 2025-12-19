package trees

type BinaryTree struct {
	Root *TreeNode[int]
}

func MaximumHeightOfTree(root *TreeNode[int]) (int, error) {
	if root == nil {
		return 0, ErrEmptyTree
	}

	leftHeight, _ := MaximumHeightOfTree(root.Left)
	rightHeight, _ := MaximumHeightOfTree(root.Right)

	return max(leftHeight, rightHeight) + 1, nil
}
