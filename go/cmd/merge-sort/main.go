package main

import (
	"fmt"

	mergesort "github.com/recurrsed/algos/internal/merge-sort"
)

func main() {
	data := []int{10, 80, 30, 90, 40}

	result := mergesort.Sort(data)

	fmt.Println(result)
}
