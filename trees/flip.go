package trees

func FlipTreeClockwise(root *TreeNode[int]) (*TreeNode[int], error) {
	// Base cases
	if root == nil {
		return root, nil
	}
	if root.Left == nil && root.Right == nil {
		return root, nil
	}

	// Recursively call the same method
	flippedRoot, err := FlipTreeClockwise(root.Left)
	if err != nil {
		return nil, err
	}

	// Rearranging main root Node after returning
	// from recursive call
	root.Left.Left = root.Right
	root.Left.Right = root
	root.Left = nil
	root.Right = nil

	return flippedRoot, nil
}
