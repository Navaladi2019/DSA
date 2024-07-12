package tree

import queue "github.com/Navaladi2019/GoRefresher/geeksDSA/Queue"

func MaxWidth(n *Node[int]) int {

	if n == nil {
		return 0
	}

	q := queue.ArrQueue[*Node[int]]{}
	q.Enqueue(n)

	count := 0
	for !q.IsEmpty() {
		s := q.Size()
		for i := 0; i < s; i++ {
			qn, _ := q.Dequeue()

			if qn.left != nil {
				q.Enqueue(qn.left)
			}
			if qn.right != nil {
				q.Enqueue(qn.right)
			}
		}
		count = max(count, s)
	}

	return count
}
