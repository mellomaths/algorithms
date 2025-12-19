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
