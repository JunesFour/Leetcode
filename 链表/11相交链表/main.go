package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 都为空，直接返回nil
	if headA == nil || headB == nil {
		return nil
	}
	// 分别遍历headA和headB
	// 如果p1为nil，就让p1等于headA
	// 如果p2为nil，就让p2等于headB
	// 以{1,2,3,5,6,7}和{4,5,6,7}举例，当对较短链表进行第二轮遍历时，必出答案，这个时候会返回5节点
	// 以{1,2,3}和{4,5,6,7}举例，当对较短链表进行第二轮遍历时，必出答案，这个时候会返回nil
	// 上述我们在做的其实是在消除两个链表的长度差
	// 不管上述两个链表相差多少，p1(len(headA)+len(headB))和p2(len(headB)+len(headA))走的路程实际上是一样的
	// 所以两轮遍历之后要么就返回相交起始节点，要么就返回nil
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

func main() {

}

// 1 2 3 4
// 5 6 7
