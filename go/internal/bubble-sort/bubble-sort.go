package bubblesort

func Sort(data []int) []int {
	n := len(data)

	for i := 0; i < n; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp

				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return data
}

func DumbSort(data []int) []int {
	stop := len(data) - 1
	i := 0

	for i <= stop {
		if data[i] > data[i+1] {
			temp := data[i]
			data[i] = data[i+1]
			data[i+1] = temp
		}
		i++

		if i == stop {
			i = 0
			stop--
		}
	}

	return data
}
