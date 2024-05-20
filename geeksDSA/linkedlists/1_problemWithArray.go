package linkedlists

/*
 Here if we see that to implement it we are shifting all the array items
 which is not good way and its time consuming
*/
func ImplementRoundRobinAlgorithm() {

	arr := [5]int{1, 2, 3, 4, 5}

	executeRoundRobin := func() {

		temp := arr[len(arr)-1]

		arr[len(arr)-1] = arr[0]

		for i := len(arr) - 2; i >= 0; i-- {

			arr[i], temp = temp, arr[i]
		}
	}

	executeRoundRobin() //becomes [2,3,4,5,1]
	executeRoundRobin() //becomes [3,4,5,1,2]
}
