package bst

// two ways to fins do inorder traversal insert into array,
// fom array you can see that its sorted due to inorder and find
// second way use dictionary to find

func CheckSumWIthTwoPairs(root *Node[int], sum int, maps map[int]struct{}) bool {
	if root == nil {
		return false
	}
	l := CheckSumWIthTwoPairs(root.Left, sum, maps)
	if l == true {
		return true
	}
	if _, ok := maps[sum-root.Value]; ok {
		return true
	}
	maps[root.Value] = struct{}{}
	return CheckSumWIthTwoPairs(root.Right, sum, maps)
}
