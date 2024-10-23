package graph

func DetectLoopRecDirectedGraph(arr [][]int, visited []bool, resStack []bool, u int) bool {
	resStack[u] = true
	visited[u] = true
	for _, v := range arr[u] {
		if resStack[v] == true {
			return true
		}
		if visited[v] == false {
			if DetectLoopRecDirectedGraph(arr, visited, resStack, v) {
				return true
			}
		}
	}

	return false
}

func DetectLoopDirectedGrapg(arr [][]int) {
	visited := make([]bool, len(arr))
	for i, _ := range arr {
		if visited[i] == false {
			recStack := make([]bool, len(arr))
			DetectLoopRecDirectedGraph(arr, visited, recStack, i)
		}
	}
}
