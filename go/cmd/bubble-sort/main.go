package main

import (
	"fmt"

	bubblesort "github.com/recurrsed/algos/internal/bubble-sort"
)

func main() {
	data := []int{12, 3, 56, 34, 5, 8}

	res := bubblesort.Sort(data)

	fmt.Println(res)
}
