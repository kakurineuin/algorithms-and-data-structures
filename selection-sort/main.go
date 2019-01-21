package main

import (
	"fmt"
)

func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		lowest := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[lowest] {
				lowest = j
			}
		}

		if i != lowest {
			arr[i], arr[lowest] = arr[lowest], arr[i]
		}
	}
}

func main() {
	arr := []int{10, 2, 6, 9, 7, 3, 1, 5, 4, 8}
	fmt.Println(arr)
	SelectionSort(arr)
	fmt.Println(arr)
}
