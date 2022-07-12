package main

import "fmt"

func main() {
	pivot := findPivot([]int{0, 1, 2, 3, 4, 5})
	fmt.Println(pivot)

	pivot = findPivot([]int{1, 2, 3, 4, 5, 0})
	fmt.Println(pivot)

	pivot = findPivot([]int{2, 3, 4, 5, 0, 1})
	fmt.Println(pivot)

	pivot = findPivot([]int{3, 4, 5, 0, 1, 2})
	fmt.Println(pivot)

	pivot = findPivot([]int{4, 5, 0, 1, 2, 3})
	fmt.Println(pivot)
	pivot = findPivot([]int{5, 0, 1, 2, 3, 4})
	fmt.Println(pivot)
	pivot = findPivot([]int{4, 1, 2, 3, 3})
	fmt.Println(pivot)
}

func findPivot(nums []int) int {
	return findPivotUtil(nums, 0, len(nums)-1)
}

func findPivotUtil(nums []int, start, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if mid+1 <= end && nums[mid] > nums[mid+1] {
		return mid + 1
	}

	if mid-1 >= start && nums[mid] < nums[mid-1] {
		return mid
	}

	if nums[mid] < nums[start] {
		return findPivotUtil(nums, start, mid-1)
	}

	return findPivotUtil(nums, mid+1, end)

}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if nums[mid] == target {
		return mid
	}

	if target < nums[mid] {
		return binarySearch(nums, start, mid-1, target)
	} else {
		return binarySearch(nums, mid+1, end, target)
	}

}
