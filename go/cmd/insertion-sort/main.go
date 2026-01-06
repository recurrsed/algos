package main

import (
	"fmt"

	insertionsort "github.com/recurrsed/algos/internal/insertion-sort"
)

func main() {
	data := []int{55, 1, 35, 56, 7, 4, 2, 1}

	res := insertionsort.Sort(data)

	fmt.Println(res)
}
