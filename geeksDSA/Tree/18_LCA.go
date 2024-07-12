package tree

func FindLCA_1(n *Node[int], n1 int, n2 int) *Node[int] {

	if n == nil {
		return n
	}

	s1, s2 := make([]*Node[int], 0, 10), make([]*Node[int], 0, 10)

	if !FindPath(n, &s1, n1) || !FindPath(n, &s2, n2) {
		return nil
	}

	for i := 0; i < min(len(s1), len(s2)); i++ {

		if i+1 >= len(s1) {
			return s1[i]
		}

		if i+1 >= len(s2) {
			return s1[i]
		}
		if s1[i+1] != s2[i+1] {
			{
				return s1[i]
			}

		}
	}

	return nil
}

func FindPath(n *Node[int], s *[]*Node[int], num int) bool {
	if n == nil {
		return false
	}

	*s = append(*s, n)

	if n.data == num {
		return true
	}

	if FindPath(n.left, s, num) || FindPath(n.right, s, num) {
		return true
	}
	*s = (*s)[0 : len(*s)-1]
	return false
}

func FindLCA_Effficient(n *Node[int], n1 int, n2 int) *Node[int] {

	if n == nil || n.data == n1 || n.data == n2 {
		return n
	}

	LLca := FindLCA_Effficient(n.left, n1, n2)
	RLca := FindLCA_Effficient(n.left, n1, n2)

	if LLca != nil && RLca != nil {
		return n
	}

	if LLca != nil {
		return LLca
	}

	return RLca

}
