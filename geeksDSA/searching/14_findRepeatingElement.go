package searching

/* here we have a constraint all elements are present between 0 and max(arr)
This is a good solution but we might have integer overflow so this might not work good
*/
func FindRepeatingElement(arr []int) int {

	maxEle := arr[0]

	totalSum := arr[0]
	for i := 1; i < len(arr); i++ {
		maxEle = max(maxEle, arr[i])
		totalSum = totalSum + arr[i]
	}

	NonDuplicateSum := maxEle * (maxEle + 1) / 2

	diffsum := totalSum - NonDuplicateSum
	difflength := len(arr) - (maxEle + 1)

	return diffsum / difflength

}

/*
This algorithm is used in linked list
*/
func FindRepeatingElement_Efficient(arr []int) int {

	slow := arr[0] + 1
	fast := arr[0] + 1

	for {

		slow = arr[slow] + 1
		fast = arr[arr[fast]+1] + 1

		if slow == fast {
			break
		}
	}

	slow = arr[0] + 1

	for slow != fast {
		slow = arr[slow] + 1
		fast = arr[fast] + 1
	}

	return slow - 1
}
