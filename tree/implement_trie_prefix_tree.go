package tree

type Trie struct {
	children [26]*Trie
	isLeaf   bool
}

func Constructor2() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	t := this
	for _, v := range word {
		c := v - 'a'
		if t.children[c] == nil {
			t.children[c] = &Trie{}
		}
		t = t.children[c]
	}
	t.isLeaf = true
}

func (this *Trie) searchPrefix(prefix string) *Trie {
	t := this
	for _, v := range prefix {
		c := v - 'a'
		if t.children[c] == nil {
			return nil
		}
		t = t.children[c]
	}
	return t
}

func (this *Trie) Search(word string) bool {
	t := this.searchPrefix(word)
	return t != nil && t.isLeaf
}

func (this *Trie) StartsWith(prefix string) bool {
	t := this.searchPrefix(prefix)
	return t != nil
}
