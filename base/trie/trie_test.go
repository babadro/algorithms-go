package trie

import (
	"log"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	log.Println(trie.Search("apple"))
	log.Println(trie.Search("app"))
	log.Println(trie.StartsWith("app"))
	trie.Insert("app")
	log.Println(trie.Search("app"))
}
