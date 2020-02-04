package kit

type MaxPQ struct {
	item []int
	N    int
}

func NewMaxPQ() *MaxPQ {
	//不使用切片的第一个位置
	return &MaxPQ{
		item: []int{-1},
		N:    0,
	}
}

func NewMaxPQWithData(data []int) *MaxPQ {
	tmp := NewMaxPQ()
	tmp.item = append(tmp.item, data...)
	return tmp
}

func (this *MaxPQ) IsEmpty() bool {
	return this.N == 0
}

func (this *MaxPQ) Size() int {
	return this.N
}

func (this *MaxPQ) Insert(key int) {
	this.item = append(this.item, key)
	this.N++
	this.swim(this.N)
}

//数值较大的会往上浮
func (this *MaxPQ) swim(k int) {
	for k > 1 && this.less(k/2, k) {
		this.exch(k/2, k)
		k = k >> 1
	}
}

//数值较小的往下沉
func (this *MaxPQ) skin(k int, n int) {
	for 2*k <= n {
		j := 2 * k
		if j < n && this.less(j, j+1) {
			j++
		}
		if this.less(k, j) {
			this.exch(k, j)
			k = j
		} else {
			break
		}
	}
}

func (this *MaxPQ) less(a, b int) bool {
	return this.item[a] < this.item[b]
}

func (this *MaxPQ) exch(a, b int) {
	this.item[a], this.item[b] = this.item[b], this.item[a]
}

func (this *MaxPQ) DeleteMax() int {
	val := this.item[1]
	this.exch(1, this.N)
	this.item = this.item[:len(this.item)-1]
	this.skin(1, this.N)
	return val
}

func (this *MaxPQ) sort() {
	n := this.N
	for i := n / 2; i >= 1; i-- {
		this.skin(i, n)
	}
	for n > 1 {
		this.exch(1, n)
		n--
		this.skin(1, n)
	}
}
