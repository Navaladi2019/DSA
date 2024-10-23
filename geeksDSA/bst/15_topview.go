package bst

import "fmt"

var mintopview, maxtopView int

type TopViewStruct struct {
	node     *Node[int]
	position int
}

func GetTopViewOfBinaryTree(n *Node[int], m map[int]int) {

	arr := make([]TopViewStruct, 0, 10)

	arr = append(arr, TopViewStruct{n, 0})

	for len(arr) > 0 {

		length := len(arr)

		for i := -0; i < length; i++ {
			ele := arr[0]
			arr = arr[1:]

			if _, ok := m[ele.position]; !ok {
				m[ele.position] = ele.node.Value
			}

			if ele.node.Left != nil {
				mintopview = min(mintopview, ele.position-1)
				arr = append(arr, TopViewStruct{ele.node.Left, ele.position - 1})
			}
			if ele.node.Right != nil {
				maxtopView = max(maxtopView, ele.position+1)
				arr = append(arr, TopViewStruct{ele.node.Left, ele.position + 1})
			}
		}
	}

}
func GetBottomViewOfBinaryTree(n *Node[int], m map[int]int) {

	arr := make([]TopViewStruct, 0, 10)

	arr = append(arr, TopViewStruct{n, 0})

	for len(arr) > 0 {

		length := len(arr)

		for i := -0; i < length; i++ {
			ele := arr[0]
			arr = arr[1:]
			m[ele.position] = ele.node.Value

			if ele.node.Left != nil {
				mintopview = min(mintopview, ele.position-1)
				arr = append(arr, TopViewStruct{ele.node.Left, ele.position - 1})
			}
			if ele.node.Right != nil {
				maxtopView = max(maxtopView, ele.position+1)
				arr = append(arr, TopViewStruct{ele.node.Left, ele.position + 1})
			}
		}
	}

}

func PrintTopViewofBinaryTree(n *Node[int]) {

	m := make(map[int]int)
	GetTopViewOfBinaryTree(n, m)

	for mintopview < maxtopView {
		fmt.Println(m[mintopview])
		mintopview++
	}
}
