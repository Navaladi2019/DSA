package graph

import (
	"fmt"
	"math"
)

type Graph [][]int

func DijkstraAlgo(graph Graph) {

	isFinalized := make([]bool, len(graph))

	distance := make([]int, len(graph))

	for i := range distance {
		distance[i] = math.MaxInt
	}

	distance[0] = 0

	for i := 0; i < len(graph); i++ {

		minindex := -1
		mindiatance := math.MaxInt
		for j := 0; j < len(isFinalized); j++ {
			if isFinalized[j] == false && distance[j] < mindiatance {
				mindiatance = distance[j]
				minindex = j
			}
		}

		for k := 0; k < len(graph); k++ {
			if graph[minindex][k] > 0 && distance[k] > distance[minindex]+graph[minindex][k] {
				if isFinalized[k] == false {
					distance[k] = distance[minindex] + graph[minindex][k]
				}
			}
		}

		isFinalized[minindex] = true

	}

	fmt.Println(distance)
}
