package 剑指Offer2

type MaxQueue struct {
	items []int
	max   []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		items: []int{},
		max:   []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.items = append(this.items, value)
	length := len(this.max)
	if length == 0 {
		this.max = append(this.max, value)
		return
	}

	if value > this.max[0] {
		this.max = this.max[0:0]
		this.max = append(this.max, value)
	} else {
		for i := len(this.max) - 1; this.max[i] < value; i-- {
			this.max = this.max[:i]
		}
		this.max = append(this.max, value)
	}
}

func (this *MaxQueue) Pop_front() int {
	if len(this.items) == 0 {
		return -1
	}
	val := this.items[0]
	this.items = this.items[1:]

	if val == this.max[0] {
		this.max = this.max[1:]
	}
	return val
}
