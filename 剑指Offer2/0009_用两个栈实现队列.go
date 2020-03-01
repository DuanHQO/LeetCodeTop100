package 剑指Offer2

type CQueue struct {
	items []int
	size  int
}

func Constructor0009() CQueue {
	return CQueue{
		items: []int{},
		size:  0,
	}
}

func (this *CQueue) AppendTail(value int) {
	this.items = append(this.items, value)
	this.size++
}

func (this *CQueue) DeleteHead() int {
	if this.size == 0 {
		return -1
	}
	val := this.items[0]
	this.items = this.items[1:]
	this.size--
	return val
}
