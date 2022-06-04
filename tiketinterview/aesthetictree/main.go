package solution

func Solution(A []int) int {
	isAesthetic := func(a []int) bool {
		for i := 1; i < len(a)-1; i++ {
			if (a[i] >= a[i-1] && a[i] <= a[i+1]) || (a[i] <= a[i-1] && a[i] >= a[i+1]) {
				return false
			}
		}
		return true
	}

	copyArr := func(originalArr []int, startIdx, endIdx int) []int {
		var out []int
		for i := startIdx; i <= endIdx; i++ {
			out = append(out, originalArr[i])
		}
		return out
	}

	var cutWaysCount int
	if isAesthetic(A) {
		// return 0 if given trees already aesthetic.
		return 0
	}

	// newT := copyArr(A, 0, len(A)-1)
	for i := 0; i < len(A)-1; i++ {
		newTrees := copyArr(A, 0, len(A)-1)
		newTrees = append(newTrees[:i], newTrees[i+1:]...)
		if isAesthetic(newTrees) {
			cutWaysCount++
		}
	}
	if cutWaysCount == 0 {
		// return -1 if there's no ways to make the given trees aesthetic.
		return -1
	}
	return cutWaysCount
}
