package main

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	// 快指针的速度更快，所以优先判断快指针是否出界
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 如果快慢指针第一次相遇，就将慢指针放到开始的位置
		// 因为len(起点到环形入口)等于len(相遇点到环形入口)
		// 所以他们第二次相遇的位置一定在环形入口
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

func main() {

}
