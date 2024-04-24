package sorting

func CycleSortDistinct(arr []int) {

	for cs := 0; cs < len(arr); cs++ {

		item := arr[cs]

		position := cs

		for j := cs + 1; j < len(arr); j++ {
			if arr[j] < item {
				position++
			}
		}

		temp := arr[position]

		arr[position] = item
		item = temp

		for position != cs {

			position = cs
			for i := cs + 1; i < len(arr); i++ {

				if arr[i] < item {
					position++
				}
			}

			temp := arr[position]

			arr[position] = item
			item = temp

		}

	}
}

func CycleSortDuplicate(arr []int) {

	for cs := 0; cs < len(arr); cs++ {

		item := arr[cs]

		position := cs

		for j := cs + 1; j < len(arr); j++ {
			if arr[j] < item {
				position++
			}
		}

		temp := arr[position]

		arr[position] = item
		item = temp

		for position != cs {

			position = cs
			for i := cs + 1; i < len(arr); i++ {

				if arr[i] <= item {
					position++
				}
			}

			temp := arr[position]

			arr[position] = item
			item = temp

		}

	}
}
