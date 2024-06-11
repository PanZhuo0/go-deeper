package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先
// Recursion --> TimeLimitExceed
func min1(root *TreeNode) int {
	// if nil Tree
	if root == nil {
		return 0
	}
	// if leaf node
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minDep := 9999
	if root.Left != nil {
		// min depth of left tree
		if min1(root.Left) < minDep {
			minDep = min1(root.Left)
		}
	}
	if root.Right != nil {
		// min depth of right tree
		if min1(root.Right) < minDep {
			minDep = min1(root.Right)
		}
	}
	// min depth of this tree
	return minDep + 1
}

// 广度优先
// 二叉树层序遍历(类比图的深度遍历)
func min(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	size := 1 //number of node of each layer
	count := 1
	for len(queue) != 0 || size != 0 {
		for size > 0 {
			// 遍历该层所有元素
			root = queue[0]
			queue = queue[1:len(queue)]
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
			if root.Left == nil && root.Right == nil {
				return count
			}
			size--
		}
		size = len(queue)
		count++
	}
	return count
}
