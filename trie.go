package trie

import (
	"log"
	"strings"
	"sync"
)

type trieNode struct {
	meta     interface{}
	children map[string]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		meta:     nil,
		children: map[string]*trieNode{},
	}
}

type Trie struct {
	// root node of the trie
	root *trieNode
	// separator between words
	sep        string
	workingMap map[string]*trieNode
	mu         sync.RWMutex
	// keysCache is a map from a string to a slice of strings. ["a.b.c:[a,b,c]"]
	keyCache   map[string][]string
	keyCacheMu sync.Mutex
}

var _ Trier = (*Trie)(nil)

func NewTrie(sep string) *Trie {
	if len(sep) > 1 {
		log.Fatal("[ERROR] not support separator len > 1")
	}
	return &Trie{
		root:       newTrieNode(),
		sep:        sep,
		workingMap: make(map[string]*trieNode, defaultCap),
		mu:         sync.RWMutex{},
		keyCache:   make(map[string][]string, defaultCap),
		keyCacheMu: sync.Mutex{},
	}

}

func (t *Trie) Get(key string) interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if node, ok := t.workingMap[key]; !ok {
		return nil
	} else {
		return node.meta
	}
}

func (t *Trie) Set(key string, meta interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.root
	for _, v := range t.getPrefix(key) {
		if len(v) == 0 {
			return
		}
		child, ok := node.children[v]
		// insert when no node under it
		if !ok {
			child = newTrieNode()
			node.children[v] = child
		}
		node = child
	}
	node.meta = meta
	t.workingMap[key] = node
}

func (t *Trie) Delete(key string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.getNode(key)
	if node != nil {
		node.meta = nil
	}
	delete(t.workingMap, key)
}

func (t *Trie) Keys() []string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var keys []string
	for k := range t.workingMap {
		keys = append(keys, k)
	}
	return keys
}

func (t *Trie) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.root = newTrieNode()
	t.keyCacheMu.Lock()
	t.keyCache = make(map[string][]string, defaultCap)
	t.keyCacheMu.Unlock()
	t.workingMap = make(map[string]*trieNode, defaultCap)
}

func (t *Trie) MatchByPrefix(prefix string) map[string]interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()

	result := map[string]interface{}{}
	var travel func(name string, node *trieNode)
	travel = func(name string, node *trieNode) {
		if node == nil {
			return
		}

		if len(name) != 0 && node.meta != nil {
			result[name] = node.meta
		}

		for k, node := range node.children {
			if len(name) != 0 {
				travel(name+t.sep+k, node)
			} else {
				travel(k, node)
			}
		}
	}

	if len(prefix) == 0 {
		travel(prefix, t.root)
	} else {
		travel(prefix, t.getNode(prefix))
	}

	return result
}

// getPrefix get and set keys from cache
func (t *Trie) getPrefix(key string) []string {
	t.keyCacheMu.Lock()
	defer t.keyCacheMu.Unlock()

	if keys, ok := t.keyCache[key]; ok {
		return keys
	} else {
		keys := strings.Split(key, t.sep)
		t.keyCache[key] = keys
		return keys
	}
}

func (t *Trie) getNode(key string) *trieNode {
	node := t.root
	for _, v := range t.getPrefix(key) {
		if len(v) == 0 {
			return nil
		}
		child, ok := node.children[v]
		node = child
		if !ok || node == nil {
			break
		}
	}
	return node
}
