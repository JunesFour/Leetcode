package main

type Trie struct {
	children [26]*Trie
	isTail   bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, v := range word {
		v = v - 'a'
		if node.children[v] == nil {
			node.children[v] = &Trie{}
		}
		node = node.children[v]
	}
	node.isTail = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, v := range prefix {
		v = v - 'a'
		if node.children[v] == nil {
			return nil
		}
		node = node.children[v]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isTail
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

func main() {

}
