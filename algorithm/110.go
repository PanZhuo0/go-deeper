package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n4 := &TreeNode{4, nil, nil}
	n44 := &TreeNode{4, nil, nil}
	n3 := &TreeNode{3, n4, n44}
	n33 := &TreeNode{3, nil, nil}
	n2 := &TreeNode{2, n3, n33}
	n22 := &TreeNode{2, nil, nil}
	n1 := &TreeNode{1, n2, n22}
	fmt.Println(isBalanced(n1))
}

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

// Helper 返回该节点的层数
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lH := height(root.Left)
	rH := height(root.Left)

	if lH == -1 || rH == -1 || abs(lH-rH) > 1 {
		// 不是平衡二叉树
		return -1
	}
	if lH > rH {
		return lH + 1
	} else {
		return rH + 1
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
