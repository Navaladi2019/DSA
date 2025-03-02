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

		// here i can get min distance and increase by o(log v) by keeping it in minheap

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

func DijkstraAlgo1(graph Graph, s int) {

	finalized := make([]bool, len(graph))

	distance := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		if i != s {
			distance[i] = math.MaxInt
		}
	}

	distance[s] = 0

	for i := 0; i < len(distance); i++ {

		indextobeFinalized := -1
		mindistance := math.MaxInt
		for j := 0; j < len(graph); j++ {
			if finalized[j] == false && distance[j] < mindistance {
				indextobeFinalized = j
				mindistance = distance[j]
			}
		}
		finalized[indextobeFinalized] = true
		for k := 0; k < len(graph); k++ {
			if finalized[k] == false && graph[indextobeFinalized][k] > 0 {
				distance[k] = min(distance[k], distance[indextobeFinalized]+graph[indextobeFinalized][k])
			}
		}
	}

	fmt.Println(distance)

}
