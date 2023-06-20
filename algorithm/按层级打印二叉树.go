package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// 定义结果列表和队列
	var res [][]int
	var queue []*TreeNode

	if root == nil {
		return res
	}

	// 将根节点加入队列
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue) // 当前层级的节点数
		var level []int

		// 处理当前层级的节点
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 将当前层级的节点值加入结果列表
		res = append(res, level)
	}

	return res
}
