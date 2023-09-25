package trie

import "testing"

var triee = NewTrie(".")

func TestPutAndGet(t *testing.T) {
	t.Log(triee.Get("124"))
	triee.Put("124", "456789")
	t.Log(triee.Get("124"))

	t.Log(triee.Get("1.2.3"))
	triee.Put("1.2.3", "000")
	t.Log(triee.Get("1.2.3"))

	t.Log(triee.Get("1/2/3"))
	triee.Put("1/2/3", "111")
	t.Log(triee.Get("1/2/3"))
}
