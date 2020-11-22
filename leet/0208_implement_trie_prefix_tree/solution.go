package _208_implement_trie_prefix_tree

type Trie2 struct {
	children [26]*Trie2
	isWord   bool
}

/** Initialize your data structure here. */
func Constructor2() Trie2 {
	return Trie2{}
}

/** Inserts a word into the trie. */
func (this *Trie2) Insert(word string) {
	trie := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		var child *Trie2
		if child = trie.children[c]; child == nil {
			child = &Trie2{}
			trie.children[c] = child
		}
		trie = child
	}
	trie.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie2) Search(word string) bool {
	trie := this
	for i := 0; i < len(word); i++ {
		if trie = trie.children[word[i]-'a']; trie == nil {
			return false
		}
	}
	return trie.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie2) StartsWith(prefix string) bool {
	trie := this
	for i := 0; i < len(prefix); i++ {
		if trie = trie.children[prefix[i]-'a']; trie == nil {
			return false
		}
	}
	return true
}
