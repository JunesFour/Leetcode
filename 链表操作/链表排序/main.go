package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMid(head)
	one, two := head, mid.Next
	// 关键操作，将链表拆分成两段
	mid.Next = nil
	// 对两段进行递归排序
	L1 := sortList(one)
	L2 := sortList(two)
	// 合并两个有序链表
	merged := Merged(L1, L2)
	return merged
}

// 使用快慢指针寻找链表中点
func findMid(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 合并两个有序链表
func Merged(L1, L2 *ListNode) *ListNode {
	L := &ListNode{}
	curr := L
	var next *ListNode
	for L1 != nil && L2 != nil {
		if L1.Val < L2.Val {
			next = L1.Next
			curr.Next = L1
			L1.Next = nil
			L1 = next
		} else {
			next = L2.Next
			curr.Next = L2
			L2.Next = nil
			L2 = next
		}
		curr = curr.Next
	}
	if L1 != nil {
		curr.Next = L1
	}
	if L2 != nil {
		curr.Next = L2
	}
	return L.Next
}

func main() {

}
