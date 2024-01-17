package main

import (
	"fmt"

	"github.com/blueambertech/sorting-demo/sorting"
)

func main() {
	unsorted := []int{3, 8, 1, 3, 6, 7, 3, 5, 6, 7, 6, 2}

	fmt.Printf("Unsorted: %v\n", unsorted)

	c := make(chan map[string][]int, 2)
	defer close(c)

	go func() {
		c <- map[string][]int{"qsort": sorting.QuickSort(unsorted, sorting.Ascending)}
	}()
	go func() {
		c <- map[string][]int{"bubblesort": sorting.BubbleSort(unsorted, sorting.Ascending)}
	}()

	fmt.Println(formatMsg(<-c))
	fmt.Println(formatMsg(<-c))
}

func formatMsg(m map[string][]int) string {
	var s string
	for v := range m {
		s += fmt.Sprintf("%s: %v", v, m[v])
	}
	return s
}
