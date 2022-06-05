// Charles and the Necklace
// Charles wants to buy a necklace in which.
// 1. There is a minimum of 1 pearl and maximum of X pearls, such that each
// pearl has its own magnificent coefficient
// 2. The pearls should be in non-decreasing order of their magnificence power.
// You are given the maximum number of pearls in a necklace and the range of
// the magnificent coefficients of the pearls. Find the number of necklaces that
// can be made that follow the mentioned conditions.​

// Input: N = 3, L = 6, R = 9
// Output: 34
// Explanation:
// The necklace can be formed in the following ways:

// The necklaces of length one that can be formed are { “6”, “7”, “8”, “9” }.
// The necklaces of length two, that can be formed are { “66”, “67”, “68”, “69”, “77”, “78”, “79”, “88”, “89”, “99” }.
// The necklaces of length three, that can be formed are { “666”, “667”, “668”, “669”, “677”, “678”, “679”, “688”, “689”, “699”, “777”, “778”, “779”, “788”, “789”, “799”, “888”, “889”, “899”, “999” }.
// Thus, in total, the necklace can be formed in (4+10+20 = 34 ) ways.

package main

import "fmt"

func main() {
	N := 3
	L := 6
	R := 9

	// Initialize 2D Array.
	var dp [][]int
	for i := 0; i < N; i++ {
		temp := make([]int, N*(R-L+1))
		dp = append(dp, temp)
	}

	count := 0
	for i := 0; i < N; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = dp[0][i-1] + 1
	}
	count = dp[0][R-L]
	for i := 1; i < N; i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
		count += dp[i][R-L]
	}
	fmt.Println(count)
}
