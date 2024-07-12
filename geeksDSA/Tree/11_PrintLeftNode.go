package tree

import (
	"fmt"

	queue "github.com/Navaladi2019/GoRefresher/geeksDSA/Queue"
)

func PrintLeftNode(n *Node[int], isLeft bool) {

	if n == nil {
		return
	}

	if isLeft {
		fmt.Println(n.data)
	}

	PrintLeftNode(n.left, true)
	PrintLeftNode(n.right, false)
}

var maxlevel int

func PrintLeftMostNode(n *Node[int], level int) {

	if n == nil {
		return
	}

	if level < maxlevel {
		fmt.Println(n.data)
	}

	PrintLeftMostNode(n.left, level+1)
	PrintLeftMostNode(n.right, level+1)
}

func PrintLeftMostNode_NewLineLeverOrderLogic(n *Node[int]) {

	q := queue.ArrQueue[*Node[int]]{}
	q.Init(10)
	q.Enqueue(n)

	for q.IsEmpty() == false {
		c := q.Size()
		for i := 0; i < c; i++ {
			qn, _ := q.Dequeue()
			if i == 0 {
				fmt.Print(qn.data, " ")
			}
			if qn.left != nil {
				q.Enqueue(qn.left)
			}
			if qn.right != nil {
				q.Enqueue(qn.right)
			}
		}
		fmt.Println()
	}
}
