
# Trie

Trie implemented by Go.



## install

To install this project run:

```bash
  go get github.com/hellolib/trie
```


## Demo

```go
package main

import (
	"fmt"
	"github.com/hellolib/trie"
)

func main() {
	var tree trie.Trier
	tree = trie.NewTrie(".")

	// Set
	tree.Set("a.b.c", 1)
	// Get
	fmt.Println("a.b.c: ", tree.Get("a.b.c"))
	tree.Set("a.b.c.d", 2)
	fmt.Println("a.b.c.d: ", tree.Get("a.b.c.d"))
	tree.Set("a.b.c.d.e", "abcde")
	fmt.Println("a.b.c.d.e: ", tree.Get("a.b.c.d.e"))
	tree.Set("a.b.c.d.f", true)
	fmt.Println("a.b.c.d.f: ", tree.Get("a.b.c.d.f"))

	// Keys
	fmt.Println("all keys: ", tree.Keys())
	// MatchByPrefix
	fmt.Println("MatchByPrefix(a.b.c.d) :", tree.MatchByPrefix("a.b.c.d"))
	fmt.Println("MatchByPrefix(a.b.c.d.e) :", tree.MatchByPrefix("a.b.c.d.e"))
	fmt.Println("MatchByPrefix(a.b.c.d.f) :", tree.MatchByPrefix("a.b.c.d.f"))

	// delete
	tree.Delete("a.b.c.d.e")
	fmt.Println("a.b.c.d.e: ", tree.Get("a.b.c.d.e"))
	fmt.Println("all keys: ", tree.Keys())
	// clear
	tree.Clear()
	fmt.Println("clear all keys: ", tree.Keys())
}

```


## Documentation

[Documentation](https://pkg.go.dev/github.com/hellolib/trie)

