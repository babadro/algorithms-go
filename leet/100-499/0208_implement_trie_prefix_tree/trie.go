package _208_implement_trie_prefix_tree

// tptl. passed
type node struct {
	children [26]*node
	endWord  bool
}

type Trie struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: new(node)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.root
	for i := range word {
		ch := word[i] - 'a'
		if cur.children[ch] == nil {
			cur.children[ch] = new(node)
		}

		cur = cur.children[ch]
	}

	cur.endWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.root
	for i := range word {
		ch := word[i] - 'a'

		if cur.children[ch] == nil {
			return false
		}

		cur = cur.children[ch]
	}

	return cur.endWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for i := range prefix {
		ch := prefix[i] - 'a'

		if cur.children[ch] == nil {
			return false
		}

		cur = cur.children[ch]
	}

	return true
}
