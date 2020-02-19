package algirithm

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
	val  byte
	sons [26]*Trie
	end  int
}

/** Initialize your data structure here. */
func ConstructorA() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		index := word[i] - 'a'
		if node.sons[index] == nil {
			node.sons[index] = &Trie{val: word[i]}
		}
		node = node.sons[index]
	}
	node.end++
	//fmt.Printf("%v\n", node)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		index := word[i] - 'a'
		if node.sons[index] == nil {
			return false
		}
		node = node.sons[index]
	}

	if node.end > 0 {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	size := len(prefix)
	for i := 0; i < size; i++ {
		index := prefix[i] - 'a'
		if node.sons[index] == nil {
			return false
		}
		node = node.sons[index]
	}
	return true
}
