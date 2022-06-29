package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 1
	slow, fast := head, head
	// 首先得到链表的长度
	end := head
	for end.Next != nil {
		end = end.Next
		n++
	}
	// 得到链表实际上要旋转的次数
	k %= n
	if k == 0 {
		return head
	}
	// 使用快慢指针得到链表的后k个节点
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 完成截取拼接操作
	end.Next = head
	begin := slow.Next
	slow.Next = nil
	return begin
}

func main() {

}
