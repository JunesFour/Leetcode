package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用hash表求解
func hasCycle(head *ListNode) bool {
	hashTable := make(map[*ListNode]bool)
	for head != nil {
		if hashTable[head] {
			return true
		}
		hashTable[head] = true
		head = head.Next
	}
	return false
}

// 使用快慢指针
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {

}
