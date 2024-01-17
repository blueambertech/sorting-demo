package sorting

import "cmp"

type SortOrder int

const Ascending SortOrder = 0
const Descending SortOrder = 1

// QuickSort takes a slice whose elements satify the cmp.Ordered interface and sorts it using the quicksort algorithm in
// the order specified by the order parameter
func QuickSort[T cmp.Ordered](arr []T, order SortOrder) []T {
	r := copySliceToNew(arr)
	qSort(r, 0, len(r)-1, order)
	return r
}

// BubbleSort takes a slice whose elements satify the cmp.Ordered interface and sorts it using the bubblesort algorithm in
// the order specified by the order parameter
func BubbleSort[T cmp.Ordered](arr []T, order SortOrder) []T {
	r := copySliceToNew(arr)
	n := len(r)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if order == 0 {
				if r[j] > r[j+1] {
					swap(r, j, j+1)
					swapped = true
				}
			} else {
				if r[j] < r[j+1] {
					swap(r, j, j+1)
					swapped = true
				}
			}
		}
		if !swapped {
			break
		}
	}
	return r
}

func qSort[T cmp.Ordered](arr []T, start, end int, order SortOrder) {
	if start < end {
		i := partition(arr, start, end, order)
		qSort(arr, start, i-1, order)
		qSort(arr, i+1, end, order)
	}
}

func partition[T cmp.Ordered](arr []T, start, end int, order SortOrder) int {
	pivot := arr[end]
	i := start - 1

	for j := start; j < end; j++ {
		if order == 0 {
			if arr[j] < pivot {
				i++
				swap(arr, i, j)
			}
		} else {
			if arr[j] > pivot {
				i++
				swap(arr, i, j)
			}
		}
	}
	swap(arr, i+1, end)
	return i + 1
}

func swap[T any](arr []T, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func copySliceToNew[T any](src []T) []T {
	n := len(src)
	dst := make([]T, n)
	for i := 0; i < n; i++ {
		dst[i] = src[i]
	}
	return dst
}
