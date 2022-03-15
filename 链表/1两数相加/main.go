package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	pre := new(ListNode)
	head := pre
	x, y := 0, 0
	for l1 != nil && l2 != nil {
		x = (l1.Val + l2.Val + y) % 10
		res = new(ListNode)
		res.Val = x
		y = (l1.Val + l2.Val + y) / 10
		pre.Next = res
		pre = res
		l1 = l1.Next
		l2 = l2.Next
		/*fmt.Println("x = ", x)
		fmt.Println("y = ", y)*/
	}
	for l1 != nil {
		res = new(ListNode)
		x = (l1.Val + y) % 10
		res.Val = x
		y = (l1.Val + y) / 10
		pre.Next = res
		pre = res
		l1 = l1.Next
	}
	for l2 != nil {
		res = new(ListNode)
		x = (l2.Val + y) % 10
		res.Val = x
		y = (l2.Val + y) / 10
		pre.Next = res
		pre = res
		l2 = l2.Next
	}
	if y != 0 {
		res = new(ListNode)
		x = y % 10
		res.Val = x
		pre.Next = res
		pre = res
	}
	return head.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func main() {
	var l1 *ListNode
	var l2 *ListNode
	var head1 *ListNode
	var head2 *ListNode
	x := 0
	for i := 0; i < 7; i++ {
		fmt.Scanf("%d", &x)
		if head1 == nil {
			head1 = &ListNode{Val: x}
			l1 = head1
		} else {
			l1.Next = &ListNode{Val: x}
			l1 = l1.Next
		}
	}
	for i := 0; i < 4; i++ {
		fmt.Scanf("%d", &x)
		if head2 == nil {
			head2 = &ListNode{Val: x}
			l2 = head2
		} else {
			l2.Next = &ListNode{Val: x}
			l2 = l2.Next
		}
	}
	res := addTwoNumbers(head1, head2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
