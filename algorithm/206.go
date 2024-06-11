package main

func main() {

}

func reverseList(head *ListNode) *ListNode {
	// 1.需要一个指针指向头节点
	// 2.需要一个指针
	if head == nil {
		return nil
	}
	pre := head
	curr := head.Next
	pre.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next, pre, curr = pre, curr, next
	}
	return pre
}
