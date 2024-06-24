package main

/* 删除链表中的倒数第n个节点
1.遍历法
2.stack
3.双指针
*/
func main() {

}

/* 双指针法 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head} // 哑巴节点
	left, right := dummy, head  //left 指向哑巴节点，right指向头节点
	for i := 0; i < n; i++ {
		// right 先行n步
		right = right.Next
	}
	for right != nil {
		// left、right 同行直到right走到nil
		left = left.Next
		right = right.Next
	}
	// 删除该节点
	left.Next = left.Next.Next
	return dummy.Next
}
