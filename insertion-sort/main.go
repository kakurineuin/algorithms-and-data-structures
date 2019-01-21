package main

import (
	"fmt"
)

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		currentVal := arr[i]
		j := i - 1

		for ; j >= 0 && currentVal < arr[j]; j-- {
			arr[j+1] = arr[j]
		}

		arr[j+1] = currentVal
	}
}

func main() {
	arr := []int{10, 2, 6, 9, 7, 3, 1, 5, 4, 8}
	fmt.Println(arr)
	InsertionSort(arr)
	fmt.Println(arr)
}
