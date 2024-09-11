package graph

import "testing"

func Test_BFS(t *testing.T) {
	arr := make([][]int, 5)

	arr[0] = append(arr[0], 1)
	arr[0] = append(arr[0], 2)
	arr[1] = append(arr[1], 0)
	arr[1] = append(arr[1], 2)
	arr[1] = append(arr[1], 3)
	arr[2] = append(arr[2], 0)
	arr[2] = append(arr[2], 1)
	arr[2] = append(arr[2], 3)
	arr[2] = append(arr[2], 4)
	arr[3] = append(arr[3], 1)
	arr[3] = append(arr[3], 2)
	arr[3] = append(arr[3], 4)
	arr[4] = append(arr[4], 2)
	arr[4] = append(arr[4], 3)
	BFS_Disconnected(arr)
}
