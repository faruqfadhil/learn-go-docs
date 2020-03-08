package main

import "fmt"

func main() {
	list := []int{30, 62, 53, 42, 17, 97, 91, 38}
	// fmt.Println(insertionSort(list))
	fmt.Println(selectionSort(list))
	// fmt.Println(bubbleSort(list))
	// fmt.Println(shellSort(list))
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
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[smallestIdx] > nums[j] {
				smallestIdx = j
			}
		}
		if nums[i] != nums[smallestIdx] {
			nums[i], nums[smallestIdx] = nums[smallestIdx], nums[i]
		}
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

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	leftSubArray := nums[:mid]
	rightSubArray := nums[mid:]
	return merge(mergeSort(leftSubArray), mergeSort(rightSubArray))
}

func merge(left, right []int) []int {
	size := len(left) + len(right)
	l := 0
	r := 0
	newSlice := make([]int, size)
	for i := 0; i < size; i++ {
		if l > len(left)-1 && r <= len(right)-1 {
			newSlice[i] = right[r]
			r++
		} else if r > len(right)-1 && l <= len(left)-1 {
			newSlice[i] = left[l]
			l++
		} else if right[r] < left[l] {
			newSlice[i] = right[r]
			r++
		} else {
			newSlice[i] = left[l]
			l++
		}
	}
	return newSlice
}
