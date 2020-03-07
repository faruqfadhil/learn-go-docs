package main

import "fmt"

func main() {
	list := []int{30, 62, 53, 42, 17, 97, 91, 38}
	// fmt.Println(insertionSort(list))
	// fmt.Println(selectionSort(list))
	// fmt.Println(bubbleSort(list))
	fmt.Println(shellSort(list))
}

func insertionSort(nums []int) []int {
	n := len(nums)
	for i := 1; i <= n-1; i++ {
		j := i
		for j > 0 {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
			j--
		}
	}

	return nums
}

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		smallestIdx := i
		for j := i; j < len(nums); j++ {
			if nums[smallestIdx] > nums[j] {
				smallestIdx = j
			}
		}
		nums[i], nums[smallestIdx] = nums[smallestIdx], nums[i]
	}

	return nums
}

func bubbleSort(nums []int) []int {
	didSwap := true
	for j := len(nums); j > 0; j-- {
		didSwap = false
		for i := 1; j > i; i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				didSwap = true
			}
		}
		if !didSwap {
			break
		}
	}
	return nums
}

func shellSort(nums []int) []int {
	jarak := len(nums)
	for jarak > 1 {
		jarak = jarak / 2
		didSwap := true
		for didSwap {
			didSwap = false
			for i := 0; i < len(nums)-jarak; i++ {
				if nums[i] > nums[i+jarak] {
					nums[i+jarak], nums[i] = nums[i], nums[i+jarak]
					didSwap = true
				}
			}

		}
	}
	return nums
}
