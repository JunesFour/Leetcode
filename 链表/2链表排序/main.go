package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMid(head)
	one, two := head, mid.Next
	mid.Next = nil
	L1 := sortList(one)
	L2 := sortList(two)
	merged := Merged(L1, L2)
	return merged
}
func findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
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
