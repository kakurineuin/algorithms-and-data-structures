package main

import (
	"fmt"
)

func pivot(arr []int, start, end int) int {
	pivot := arr[start] // 以第一個當 pivot。
	swapIndex := start

	for i := start + 1; i <= end; i++ {
		if arr[i] < pivot {
			swapIndex++
			arr[swapIndex], arr[i] = arr[i], arr[swapIndex]
		}
	}

	arr[start], arr[swapIndex] = arr[swapIndex], arr[start]
	return swapIndex
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivotIndex := pivot(arr, left, right)
		QuickSort(arr, left, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, right)
	}
}

func main() {
	arr := []int{10, 2, 6, 9, 7, 3, 1, 5, 4, 8}
	fmt.Println(arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
