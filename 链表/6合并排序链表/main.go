package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var L *ListNode
	L = mergeList(lists[0], lists[1])
	for i := 2; i < len(lists); i++ {
		L = mergeList(L, lists[i])
	}
	return L
}

func mergeList(L1, L2 *ListNode) *ListNode {
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
