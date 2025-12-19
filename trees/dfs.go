package trees

func DfsInOrder(node *TreeNode[int], values *[]int) {
	if node != nil {
		DfsInOrder(node.Left, values)
		*values = append(*values, node.Value)
		DfsInOrder(node.Right, values)
	}
}

func DfsPreOrder(node *TreeNode[int], values *[]int) {
	if node != nil {
		*values = append(*values, node.Value)
		DfsPreOrder(node.Left, values)
		DfsPreOrder(node.Right, values)
	}
}

func DfsPostOrder(node *TreeNode[int], values *[]int) {
	if node != nil {
		DfsPostOrder(node.Left, values)
		DfsPostOrder(node.Right, values)
		*values = append(*values, node.Value)
	}
}
