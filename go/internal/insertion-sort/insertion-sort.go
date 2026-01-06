package insertionsort

func Sort(data []int) []int {
	for i := 1; i < len(data); i++ {
		current := data[i]
		k := i - 1

		for k >= 0 && data[k] > current {
			data[k+1] = data[k]
			k--
		}

		data[k+1] = current
	}

	return data
}
