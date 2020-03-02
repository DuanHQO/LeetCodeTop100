package 剑指Offer2

type MinStack struct {
	items []int
	min   []int
}

/** initialize your data structure here. */
func Constructor0030() MinStack {
	return MinStack{
		items: []int{},
		min:   []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.items = append(this.items, x)
	length := len(this.min)
	if length == 0 {
		this.min = append(this.min, x)
		return
	}

	if this.min[length-1] < x {
		this.min = append(this.min, this.min[length-1])
	} else {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if this.items == nil || len(this.items) == 0 {
		return
	}
	this.items = this.items[:len(this.items)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if this.items == nil || len(this.items) == 0 {
		return 0
	}
	return this.items[len(this.items)-1]
}

func (this *MinStack) Min() int {
	if this.items == nil || len(this.items) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}
