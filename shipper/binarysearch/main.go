package main

import (
	"fmt"
)

func main() {
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		var N int
		fmt.Scan(&N)
		arr := make([]int, N)
		for i := 0; i < N; i++ {
			if _, err := fmt.Scan(&arr[i]); err != nil {
				panic(err)
			}
		}
		fmt.Println(getAnswer(arr))
	}
}

func getAnswer(arr []int) string {
	out := linierSearch(arr)
	if len(out) < 1 {
		return "Not Found"
	}
	var aggsOut string
	for _, o := range out {
		aggsOut += fmt.Sprintf("%d ", o)
	}
	return aggsOut
}

func binarySearch(needle int, haystack []int) int {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return -1
	}
	return low
}

func linierSearch(arr []int) []int {
	out := []int{}
	for i, val := range arr {
		fmt.Println(val)
		if val == i+1 {
			out = append(out, i+1)
		}
	}
	return out
}

// func main() {
// 	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
// 	fmt.Println(binarySearch(20, items))
// }
