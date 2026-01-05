package main

import (
	"fmt"

	quicksort "github.com/recurrsed/algos/internal/quick-sort"
)

func main() {
	data := []int{10, 80, 30, 90, 40}
	res := quicksort.Sort(data, 0, len(data)-1)

	fmt.Println(res)
}
