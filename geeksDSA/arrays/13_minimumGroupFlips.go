package arrays

import "fmt"

func MinimumConsecutiveFlips(arr []int) (flipitems int, count int) {

	count1 := 0
	count0 := 0

	if arr[0] == 0 {
		count0++
	} else {
		count1++
	}
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == 0 && arr[i] == 1 {
			count1++
		}
		if arr[i-1] == 1 && arr[i] == 0 {
			count0++
		}
	}

	if count0 <= count1 {
		count = count0
	} else {
		count = count1
		flipitems = 1
	}

	return
}

/*

Here there is a trick , check the first
and last element if its same (1) then we need to flip 0 and vice versa
and if its different then flipping any one ( 1 or 0) is same

*/

func MinimumConsecutiveFlips_Efficient(arr []int) (res int) {

	elementtoFlip := 0
	isflipStaRTED := false

	if arr[0] == arr[len(arr)-1] {
		if arr[0] == 0 {
			elementtoFlip = 1
		}
	}

	for i := 0; i < len(arr); i++ {
		if (i == 0 || arr[i-1] != elementtoFlip) && elementtoFlip == arr[i] {
			fmt.Print("flipping from ", i, "- ")
			isflipStaRTED = true
		}

		if (arr[i] != elementtoFlip || i == len(arr)-1) && isflipStaRTED == true {
			fmt.Print(i-1, "\n")
			isflipStaRTED = false
		}

	}

	return

}
