package graph

import (
	"fmt"

	"github.com/Navaladi2019/GoRefresher/geeksDSA/stack"
)

func TopologicalSOrtDFSREC(arr [][]int, u int, visited []bool, st *stack.ArrStack[int]) {
	visited[u] = true
	for _, v := range arr[u] {
		if visited[v] == false {
			TopologicalSOrtDFSREC(arr, v, visited, st)
		}
	}
	st.Push(u)
}

func TopologicalSOrtDFS(arr [][]int) {
	visited := make([]bool, len(arr))
	sta := stack.ArrStack[int]{}
	for i, _ := range arr {

		if visited[i] == false {
			TopologicalSOrtDFSREC(arr, i, visited, &sta)
		}
	}
	for !sta.IsEmpty() {
		fmt.Println(sta.Pop())
	}
}
