package tree

import (
	"fmt"

	queue "github.com/Navaladi2019/GoRefresher/geeksDSA/Queue"
)

func LevelOrderTraversal_Naive(n *Node[int]) {
	h := height(n)
	for i := 0; i <= h; i++ {
		PrintKNode(n, i)
	}
}

func LevelOrderTraversalEfficient(n *Node[int]) {

	q := queue.ArrQueue[*Node[int]]{}
	q.Init(10)
	q.Enqueue(n)
	//q.Enqueue(nil)

	for q.IsEmpty() == false {

		qn, _ := q.Dequeue()

		fmt.Println(qn.data)
		if qn.left != nil {
			q.Enqueue(qn.left)
		}
		if qn.right != nil {
			q.Enqueue(qn.right)
		}
	}

}

func LevelOrderTraversalEfficient_NewLine_1(n *Node[int]) {

	q := queue.ArrQueue[*Node[int]]{}
	q.Init(10)
	q.Enqueue(n)
	q.Enqueue(nil)

	for q.IsEmpty() == false {

		qn, _ := q.Dequeue()

		if qn == nil {
			fmt.Println()
			if q.IsEmpty() {
				break
			}
			q.Enqueue(nil)
			continue
		}

		fmt.Print(qn.data, " ")
		if qn.left != nil {
			q.Enqueue(qn.left)
		}
		if qn.right != nil {
			q.Enqueue(qn.right)
		}
	}

}

func LevelOrderTraversalEfficient_NewLine_2(n *Node[int]) {

	q := queue.ArrQueue[*Node[int]]{}
	q.Init(10)
	q.Enqueue(n)

	for q.IsEmpty() == false {
		c := q.Size()
		for i := 0; i < c; i++ {
			qn, _ := q.Dequeue()
			fmt.Print(qn.data, " ")
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
