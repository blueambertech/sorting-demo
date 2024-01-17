package sorting

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestQuickSortAsc(t *testing.T) {
	unsorted := []int{5, 6, 1, 2, 8, 3, 7, 4, 9}
	correct := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sorted := QuickSort(unsorted, Ascending)

	if !reflect.DeepEqual(sorted, correct) {
		t.Errorf("QuickSort: Did not sort correctly, expected %v got %v", correct, sorted)
	}
}

func TestBubbleSortAsc(t *testing.T) {
	unsorted := []int{5, 3, 1, 2, 8, 6, 7, 4, 9}
	correct := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sorted := BubbleSort(unsorted, Ascending)

	if !reflect.DeepEqual(sorted, correct) {
		t.Errorf("BubbleSort: Did not sort correctly, expected %v got %v", correct, sorted)
	}
}

func TestQuickSortDesc(t *testing.T) {
	unsorted := []int{5, 6, 1, 2, 8, 3, 7, 4, 9}
	correct := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	sorted := QuickSort(unsorted, Descending)

	if !reflect.DeepEqual(sorted, correct) {
		t.Errorf("QuickSort: Did not sort correctly, expected %v got %v", correct, sorted)
	}
}

func TestBubbleSortDesc(t *testing.T) {
	unsorted := []int{5, 3, 1, 2, 8, 6, 7, 4, 9}
	correct := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	sorted := BubbleSort(unsorted, Descending)

	if !reflect.DeepEqual(sorted, correct) {
		t.Errorf("BubbleSort: Did not sort correctly, expected %v got %v", correct, sorted)
	}
}

func BenchmarkQuickSortSmall(b *testing.B) {
	unsorted := []int{}
	r := rand.New(rand.NewSource(555))
	for i := 0; i < 50; i++ {
		unsorted = append(unsorted, r.Int())
	}

	for i := 0; i < 1000; i++ {
		_ = QuickSort(unsorted, Ascending)
	}
}

func BenchmarkBubbleSortSmall(b *testing.B) {
	unsorted := []int{}
	r := rand.New(rand.NewSource(555))
	for i := 0; i < 50; i++ {
		unsorted = append(unsorted, r.Int())
	}

	for i := 0; i < 1000; i++ {
		_ = BubbleSort(unsorted, Ascending)
	}
}

func BenchmarkQuickSortLarge(b *testing.B) {
	unsorted := []int{}
	r := rand.New(rand.NewSource(555))
	for i := 0; i < 2500; i++ {
		unsorted = append(unsorted, r.Int())
	}

	for i := 0; i < 1000; i++ {
		_ = QuickSort(unsorted, Ascending)
	}
}

func BenchmarkBubbleSortLarge(b *testing.B) {
	unsorted := []int{}
	r := rand.New(rand.NewSource(555))
	for i := 0; i < 2500; i++ {
		unsorted = append(unsorted, r.Int())
	}

	for i := 0; i < 1000; i++ {
		_ = BubbleSort(unsorted, Ascending)
	}
}
