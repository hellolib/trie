package trie

type Trier interface {
	Set(key string, meta interface{})
	Get(key string) interface{}
	MatchByPrefix(prefix string) map[string]interface{}
	Delete(key string)
	Keys() []string
	Clear()
}
