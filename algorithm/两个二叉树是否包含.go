package main

import (
	"fmt"
	"strings"
)

type TwoTreeNode struct {
	Val   int
	Left  *TwoTreeNode
	Right *TwoTreeNode
}

func isSubtree(s *TwoTreeNode, t *TwoTreeNode) bool {
	sStr := generatePreorderString(s)
	tStr := generatePreorderString(t)
	return strings.Contains(sStr, tStr)
}

// 生成二叉树的前序遍历字符串
func generatePreorderString(root *TwoTreeNode) string {
	if root == nil {
		return ",null"
	}
	left := generatePreorderString(root.Left)
	right := generatePreorderString(root.Right)
	return fmt.Sprintf(",%d%s%s", root.Val, left, right)
}

//时间复杂度为O(m×n)。
