package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{-1, nil, nil}
	n2 := &TreeNode{1, n1, nil}
	n3 := &TreeNode{3, nil, nil}
	n4 := &TreeNode{-2, nil, nil}
	n5 := &TreeNode{-2, n2, nil}
	n6 := &TreeNode{-3, n3, n4}
	n7 := &TreeNode{1, n5, n6}
	fmt.Println(hasPathSum(n7, -4))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	sum := 0
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			// if root == p(origin root) do not add root.Val again
			sum += root.Val
			if sum == targetSum && root.Left == nil && root.Right == nil {
				return true
			}
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1:][0]
		stack = stack[:len(stack)-1]
		// 什么时候可以进行删除?
		if root.Right == nil {
			sum -= root.Val
		}
		root = root.Right
	}
	return false
}
