package trees

func BfsLevelOrderTraversal(root *TreeNode[int]) ([]int, error) {
	queue := []*TreeNode[int]{}
	queue = append(queue, root)
	values := []int{}
	for len(queue) > 0 {
		node := queue[0]
		values = append(values, node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	return values, nil
}

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
