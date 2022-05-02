package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome1(head *ListNode) (res bool) {
	res = true
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	i, j := 0, len(arr)-1
	for i < j {
		if arr[i] != arr[j] {
			res = false
			break
		}
		i++
		j--
	}
	return
}

func isPalindrome2(head *ListNode) bool {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	n := len(arr)
	for i, v := range arr[:n/2] {
		if v != arr[n-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	// 反转后一段链表
	var head2 *ListNode = nil
	for slow != nil {
		next := slow.Next
		slow.Next = head2
		head2 = slow
		slow = next
	}
	// 对比
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func main() {
	L := &ListNode{}
	head := L
	for i := 0; i < 5; i++ {
		var x int
		fmt.Scanf("%d", &x)
		L.Next = &ListNode{Val: x}
		L = L.Next
	}
	fmt.Println(isPalindrome2(head.Next))
	/*head = head.Next
	  for head != nil {
	       fmt.Println(head.Val)
	       head = head.Next
	  }*/
}
