package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	L := &ListNode{}
	curr := L
	var next *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			next = l1.Next
			curr.Next = l1
			l1.Next = nil
			l1 = next
		} else {
			next = l2.Next
			curr.Next = l2
			l2.Next = nil
			l2 = next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return L.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func main() {

}
