package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

func RadixSort(arr []int) []int {
	bucketMap := map[int][]int{
		0: make([]int, 0),
		1: make([]int, 0),
		2: make([]int, 0),
		3: make([]int, 0),
		4: make([]int, 0),
		5: make([]int, 0),
		6: make([]int, 0),
		7: make([]int, 0),
		8: make([]int, 0),
		9: make([]int, 0),
	}

	var num int
	mostDigits := MostDigits(arr)

	for place := 0; place < mostDigits; place++ {

		// 將 arr 中的元素依照數字分組。
		for len(arr) > 0 {
			num, arr = arr[0], arr[1:] // shift，移除並取得第一個元素。
			digit := GetDigit(num, place)
			bucketMap[digit] = append(bucketMap[digit], num)
		}

		// 依照 0 ~ 9 的順序，從各組中將元素放回 arr。
		for j := 0; j < 10; j++ {
			for len(bucketMap[j]) > 0 {
				num, bucketMap[j] = bucketMap[j][0], bucketMap[j][1:] // shift，移除並取得第一個元素。
				arr = append(arr, num)
			}
		}
	}

	return arr
}

func GetDigit(num, place int) int {
	strNum := strconv.Itoa(num)
	length := len(strNum)

	if place >= length {
		return 0
	} else {
		digit, err := strconv.Atoi(string(strNum[length-1-place]))

		if err != nil {
			panic(err)
		}

		return digit
	}
}

func DigitCount(num int) int {
	return utf8.RuneCountInString(strconv.Itoa(num))
}

func MostDigits(nums []int) int {
	maxDigits := 0.0

	for _, num := range nums {
		maxDigits = math.Max(maxDigits, float64(DigitCount(num)))
	}

	return int(maxDigits)
}

func main() {
	arr := []int{10, 2, 61, 92, 7, 352, 16, 5000, 41, 28, 1, 3, 36, 2163}
	fmt.Println(arr)
	arr = RadixSort(arr)
	fmt.Println(arr)
}
