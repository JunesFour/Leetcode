package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 移动链表
func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 1
	slow, fast := head, head
	end := head
	for end.Next != nil {
		end = end.Next
		n++
	}
	k %= n
	if k == 0 {
		return head
	}
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	end.Next = head
	begin := slow.Next
	slow.Next = nil
	return begin
}

// 不移动链表
func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var arr1 []int
	var arr2 []int
	begin := head
	for begin != nil {
		arr1 = append(arr1, begin.Val)
		begin = begin.Next
	}
	// 1 2 3 4 5 6
	length := len(arr1)
	k %= length
	if k == 0 {
		return head
	}
	for i := 0; i < length; i++ {
		arr2 = append(arr2, arr1[(length-k+i)%length])
	}
	begin = head
	for _, v := range arr2 {
		begin.Val = v
		begin = begin.Next
	}
	return head
}

// 官方解法
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	add := n - k%n
	if add == n {
		return head
	}
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}

func main() {

}
