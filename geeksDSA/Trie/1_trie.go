package trie

type Node struct {
	Value  string
	IsLeaf bool
	Child  []*Node
}

func (node *Node) Add(value string) {

	curr := node
	for _, char := range []rune(value) {

		if curr.Child == nil {
			child := make([]*Node, 26)
			curr.Child = child
		}

		index := char - 'a'

		if curr.Child[index] == nil {
			curr.Child[index] = &Node{Value: string(char)}
		}
		curr.Child[index].Value = string(char)
		curr = curr.Child[index]
	}
	curr.IsLeaf = true
}

func (node *Node) Search(value string) bool {

	curr := node
	for _, char := range []rune(value) {

		if curr == nil || curr.Child == nil {
			return false
		}
		index := char - 'a'
		curr = curr.Child[index]
	}
	return curr != nil && curr.IsLeaf == true
}
