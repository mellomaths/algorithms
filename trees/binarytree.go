package trees

type BinaryTree struct {
	Root *TreeNode[int]
}

func TreeHeight(root *TreeNode[int]) (int, error) {
	if root == nil {
		return 0, ErrEmptyTree
	}

	leftHeight, _ := TreeHeight(root.Left)
	rightHeight, _ := TreeHeight(root.Right)

	return max(leftHeight, rightHeight) + 1, nil
}
