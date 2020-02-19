package algirithm

type LRUCache struct {
	Cap  int
	Val  int
	Head *LRUNode
	Hash map[int]*LRUNode
}

type LRUNode struct {
	Key  int
	Val  int
	Last *LRUNode
	Next *LRUNode
}

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Val:  0,
		Head: nil,
		Hash: make(map[int]*LRUNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.Hash[key]
	if !ok {
		return -1
	}

	if this.Head == val {
		return val.Val
	} else if this.Val == 2 {
		this.Head = val
		return val.Val
	} else if this.Head.Last == val {
		this.Head = val
		return val.Val
	}

	//访问过后会被放在最前端
	last := val.Last
	next := val.Next
	end := this.Head.Last
	last.Next = next
	next.Last = last
	val.Last = end
	val.Next = this.Head
	end.Next = val
	this.Head.Last = val
	this.Head = val
	return val.Val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Hash[key]; ok {
		this.Hash[key].Val = value
		this.Get(key)
		return
	}

	val := &LRUNode{
		Key:  key,
		Val:  value,
		Last: nil,
		Next: nil,
	}
	this.Val++
	if this.Head == nil {
		this.Head = val
		val.Next = val
		val.Last = val
		this.Hash[key] = val
		return
	}

	end := this.Head.Last
	if this.Val > this.Cap {
		this.Val--
		delete(this.Hash, end.Key)
		end = end.Last
	}
	end.Next = val
	val.Last = end
	val.Next = this.Head
	this.Head.Last = val
	this.Head = val
	this.Hash[key] = val
}
