// prefix tree
// https://leetcode.com/problems/implement-trie-prefix-tree/
package trie

type Trie struct {
	root Node
}

type Node struct {
	isEnd bool
	links []Node
}

func (tn Node) isNull() bool {
	return tn.links == nil
}

func NewNode() Node {
	return Node{links: make([]Node, 26)}
}

func Constructor() Trie {
	return Trie{root: Node{links: make([]Node, 26)}}
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
