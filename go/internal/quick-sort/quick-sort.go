package quicksort

func Sort(data []int, start int, end int) []int {
	if start < end {
		pivot := partitionEnd(data, start, end)

		Sort(data, start, pivot-1)
		Sort(data, pivot+1, end)
	}

	return data
}

func partition(data []int, start int, end int) int {
	pivot := end
	i := start - 1
	j := start

	for j < pivot {
		if data[j] < data[pivot] {
			i++

			temp := data[i]
			data[i] = data[j]
			data[j] = temp
		}

		j++
	}

	temp := data[i+1]
	data[i+1] = data[pivot]
	data[end] = temp

	return i + 1
}

func partitionEnd(data []int, start int, end int) int {
	pivot := start
	i := end + 1
	j := end

	for j > pivot {
		if data[j] > data[pivot] {
			i--

			temp := data[i]
			data[i] = data[j]
			data[j] = temp
		}

		j--
	}

	temp := data[i-1]
	data[i-1] = data[pivot]
	data[pivot] = temp

	return i - 1
}
