package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//Tree 1
	t12 := &TreeNode{2, nil, nil}
	t13 := &TreeNode{3, nil, nil}
	t11 := &TreeNode{1, t12, t13}
	// Tree 2
	t22 := &TreeNode{2, nil, nil}
	t23 := &TreeNode{3, nil, nil}
	t21 := &TreeNode{1, t22, t23}

	fmt.Println(isSameTree(t11, t21))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	stackP := make([]*TreeNode, 0)
	stackQ := make([]*TreeNode, 0)

	if p != nil && q == nil || p == nil && q != nil {
		return false
	}

	for p != nil || len(stackP) != 0 {
		for p != nil {
			if q == nil {
				return false
			}
			if p.Val != q.Val {
				return false
			}
			stackP = append(stackP, p)
			stackQ = append(stackQ, q)

			if p.Left != nil && q.Left == nil || p.Left == nil && q.Left != nil {
				return false
			}
			p = p.Left
			q = q.Left
		}
		p = stackP[len(stackP)-1:][0]
		stackP = stackP[:len(stackP)-1]
		q = stackQ[len(stackQ)-1:][0]
		stackQ = stackQ[:len(stackQ)-1]

		if p.Right != nil && q.Right == nil || p.Right == nil && q.Right != nil {
			return false
		}
		p = p.Right
		q = q.Right
	}
	return true
}
