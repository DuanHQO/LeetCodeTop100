package algirithm

type MyStack struct {
	items []int
	size  int
}

/** Initialize your data structure here. */
func Constructor225() MyStack {
	return MyStack{
		items: []int{},
		size:  0,
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.items = append([]int{x}, this.items...)
	this.size++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	val := this.items[0]
	this.items = this.items[1:]
	this.size--
	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	val := this.items[0]
	return val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
