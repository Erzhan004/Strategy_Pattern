package main

import "fmt"

type Sorting interface {
	Sort(arr []int)
}

type bubbleSort struct {
}

func NewBubbleSort() Sorting {
	return &bubbleSort{}
}

func (b *bubbleSort) Sort(arr []int) {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				// Swap elements arr[i-1] and arr[i]
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
	}
}

type mergeSort struct {
}

func NewMergeSort() Sorting {
	return &mergeSort{}
}

func (q *mergeSort) Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	mergeSort := NewMergeSort()
	mergeSort.Sort(left)
	mergeSort.Sort(right)

	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}

}

type insertionSort struct {
}

func NewInsertionSort() Sorting {
	return &insertionSort{}
}

func (i *insertionSort) Sort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func main() {
	var choice int
	var sortMethod Sorting
	fmt.Scan(&choice)
	arr := []int{5, 6, 9, 5, 7, 2}
	switch choice {
	case 1:
		sortMethod = NewBubbleSort()
		fmt.Print("Result of Bubble Sort : ")
	case 2:
		sortMethod = NewMergeSort()
		fmt.Print("Result of Merge Sort : ")
	case 3:
		sortMethod = NewInsertionSort()
		fmt.Print("Result of Insertion Sort  : ")
	default:
		fmt.Println("Invalid choice")
		return
	}
	sortMethod.Sort(arr)
	fmt.Print(arr)
}
