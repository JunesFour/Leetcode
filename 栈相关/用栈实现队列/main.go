package main

type MyQueue struct {
	stack1, stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) != 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	front := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return front
}

func (this *MyQueue) Peek() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) != 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	front := this.stack2[len(this.stack2)-1]
	return front
}

func (this *MyQueue) Empty() bool {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return true
	} else {
		return false
	}
}

func main() {

}
