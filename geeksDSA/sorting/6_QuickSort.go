package sorting

func QuickSortUsingLomutoPartition(arr []int, low int, high int) {

REC:
	if low < high {
		p := Lomuto_partition(arr, low, high)
		QuickSortUsingLomutoPartition(arr, low, p-1)
		low = p + 1
		goto REC
		//QuickSortUsingLomutoPartition(arr, p+1, high)
	}
}

func QuickSortUsingHoaresPartition(arr []int, low int, high int) {

REC:
	if low < high {
		p := Hoares_Partition(arr, low, high)
		QuickSortUsingHoaresPartition(arr, low, p)
		low = p + 1
		//QuickSortUsingHoaresPartition(arr, p+1, high)
		goto REC
	}
}
