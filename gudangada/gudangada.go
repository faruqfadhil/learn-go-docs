package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{3, 1, 3}
	out := [][]int{}
	for i := 0; i < len(n); i++ {
		n[i] = n[i] * n[i]
	}
	sort.Ints(n)

	firtsNumIteration := 0
	secondNumIteration := len(n) - 2
	lastNumIteration := len(n) - 1

	for firtsNumIteration < secondNumIteration {
		if n[firtsNumIteration]+n[secondNumIteration] == n[lastNumIteration] {
			out = append(out, []int{n[firtsNumIteration], n[secondNumIteration], n[lastNumIteration]})
			firtsNumIteration++
			secondNumIteration--
			lastNumIteration--
		} else {

		}
	}

	fmt.Println(out)
}
