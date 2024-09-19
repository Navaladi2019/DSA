package arrays

import "fmt"

func FindFrequencyInSortedArray(arr []int) {

	frequency := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			fmt.Println(arr[i-1], ":", frequency)
			frequency = 1
		} else {
			frequency++
		}
	}

	if frequency >= 1 && len(arr) >= 1 {
		fmt.Println(arr[len(arr)-1], ":", frequency)
	}
}
