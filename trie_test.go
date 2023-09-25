package trie

import "testing"

func TestTrie(t *testing.T) {

	var triee = NewTrie(".")

	t.Log(triee.Get("124"))
	triee.Set("124", "456789")
	t.Log(triee.Get("124"))

	t.Log(triee.Get("1.2.3"))
	triee.Set("1.2.3", "000")
	t.Log(triee.Get("1.2.3"))

	t.Log(triee.Keys())

	triee.Clear()
	t.Log("clean------------------------")
	t.Log(triee.Keys())

	triee.Set("124", "456789")
	triee.Set("1.2.3", "0003")
	triee.Set("1.2.3.4", "0004")
	triee.Set("1.2.3.5", "005")
	triee.Set("1.2.3.6", "006")

	t.Log("1.2.3")
	t.Log(triee.MatchByPrefix("1.2.3"))

	t.Log("1.2.3.4")
	t.Log(triee.MatchByPrefix("1.2.3.4"))

}
