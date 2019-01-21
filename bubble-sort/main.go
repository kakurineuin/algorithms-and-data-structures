package main

import (
	"fmt"
)

func BubbleSort(arr []int) {
	for i := len(arr); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{10, 2, 6, 9, 7, 3, 1, 5, 4, 8}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)
}
