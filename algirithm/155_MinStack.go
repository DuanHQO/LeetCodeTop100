package algirithm

import (
	_ "go/types"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type MinStack struct {
	item []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		item: []int{},
		min:  []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.item = append(this.item, x)
	length := len(this.min)
	if length == 0 {
		this.min = append(this.min, x)
		return
	}
	if this.min[length-1] > x {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[length-1])
	}
}

func (this *MinStack) Pop() {
	if this.item == nil || len(this.item) == 0 {
		return
	}
	this.item = this.item[:len(this.item)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if this.item == nil || len(this.item) == 0 {
		return 0
	}
	return this.item[len(this.item)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
