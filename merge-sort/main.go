package main

import (
	"fmt"
	"math"
)

func MergeSort(arr []int) []int {
	length := len(arr)

	if length <= 1 {
		return arr
	}

	mid := int(math.Floor(float64(length) / 2))
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)
}

func Merge(arr1 []int, arr2 []int) []int {
	result := make([]int, 0)
	i := 0
	j := 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	if i < len(arr1) {
		result = append(result, arr1[i:]...)
	}

	if j < len(arr2) {
		result = append(result, arr2[j:]...)
	}

	return result
}

func main() {
	arr := []int{10, 2, 6, 9, 7, 3, 1, 5, 4, 8}
	fmt.Println(arr)
	arr = MergeSort(arr)
	fmt.Println(arr)
}
