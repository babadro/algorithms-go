package _208_implement_trie_prefix_tree

type Trie struct {
	root TrieNode
}

func Constructor() Trie {
	return Trie{root: TrieNode{links: make([]TrieNode, 26)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for i, v := range word {
		if node.links[v-'a'].isNull() {
			node.links[v-'a'] = NewNode()
		}
		if i < len(word)-1 {
			node = node.links[v-'a']
		} else {
			node.links[v-'a'].isEnd = true
		}
	}
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, v := range word {
		if node.links[v-'a'].isNull() {
			return false
		}
		node = node.links[v-'a']
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, v := range prefix {
		if node.links[v-'a'].isNull() {
			return false
		}
		node = node.links[v-'a']
	}
	return true
}

type TrieNode struct {
	isEnd bool
	links []TrieNode
}

func (tn TrieNode) isNull() bool {
	return tn.links == nil
}

func NewNode() TrieNode {
	return TrieNode{links: make([]TrieNode, 26)}
}
