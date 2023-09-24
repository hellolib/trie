package trie

import (
	"sync"
)

type trieNode struct {
	meta     interface{}
	children map[string]*trieNode
}

type Trie struct {
	// root node of the trie
	root *trieNode
	// separator between words
	sep string
	m   map[string]*trieNode
	mu  sync.RWMutex
	// cache is a map from a string to a slice of strings.
	cache   map[string][]string
	cacheMu sync.Mutex
}

func (t *Trie) Clean() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root = t.newNode()
	t.cacheMu.Lock()
	t.cache = make(map[string][]string, defaultCap)
	t.cacheMu.Unlock()
	t.m = make(map[string]*trieNode, defaultCap)
}

func (t *Trie) newNode() *trieNode {
	return &trieNode{
		meta:     nil,
		children: map[string]*trieNode{},
	}
}
