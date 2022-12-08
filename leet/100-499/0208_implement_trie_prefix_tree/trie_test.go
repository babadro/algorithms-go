package _208_implement_trie_prefix_tree

import (
	"log"
	"testing"
)

func TestTrie(t *testing.T) {
	trie3 := Constructor()
	trie3.Insert("apple")
	log.Println(trie3.Search("apple"))
	log.Println(trie3.Search("app"))
	log.Println(trie3.StartsWith("app"))
	trie3.Insert("app")
	log.Println(trie3.Search("app"))
}
