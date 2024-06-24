package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	循环迭代法

1.使用dummyHead指向头节点，便于实现交换头节点与head.Next节点
2.循环迭代

	curr:=dummyHead
	循环终止条件:
		curr.Next==nil||curr.Next.Next==nil
	节点的交换:
		temp1:=curr.Next
		temp2:=curr.Next.Next.Next
		curr.Next=curr.Next.Next
		curr.Next.Next=temp1
		temp1.Next=temp2
		curr=temp2
*/

func swapPairs(head *ListNode) *ListNode {
	// 1.new dummyHead
	dummyHead := &ListNode{0, head}
	curr := dummyHead
	// 2.循环遍历
	for curr.Next != nil && curr.Next.Next != nil {
		temp1 := curr.Next
		temp2 := curr.Next.Next.Next
		curr.Next = curr.Next.Next
		curr.Next.Next = temp1
		temp1.Next = temp2
		curr = curr.Next.Next
	}
	return dummyHead.Next
}
