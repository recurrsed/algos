package mergesort

func Sort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := data[0:mid]
	right := data[mid:]

	return merge(Sort(left), Sort(right))
}

func merge(left []int, right []int) []int {
	result := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
