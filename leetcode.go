package main

import "fmt"
import "math"

// 474. Ones and Zeroes
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func create2Darray(r int, c int, val int) [][]int {
	memo := make([][]int, r)
	for x := range memo {
		memo[x] = make([]int, c)
		for j := 0; j < c; j++ {
			memo[x][j] = val
		}
	}
	return memo
}
func findMaxForm(strs []string, m int, n int) int {
	memo := make([][]int, m+1)
	for x := range memo {
		memo[x] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		var ones, zero int = 0, 0
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '1' {
				ones++
			} else {
				zero++
			}
		}
		for j := m; j >= zero; j-- {
			for k := n; k >= ones; k-- {
				memo[j][k] = max(memo[j][k], 1+memo[j-zero][k-ones])
			}
		}
	}
	return memo[m][n]
}

// 377. Combination Sum IV
var dp []int

func combinationSum4Dp(nums []int, target int) int {
	if dp[target] >= 0 {
		return dp[target]
	}
	var total int = 0
	for i := 0; i < len(nums); i++ {
		if target >= nums[i] {
			total += combinationSum4Dp(nums, target-nums[i])
		}
	}
	dp[target] = total
	return total
}

func combinationSum4(nums []int, target int) int {
	dp = make([]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = -1
	}
	dp[0] = 1
	return combinationSum4Dp(nums, target)
}

func min(a int, b int) int { // math.Min is float type
	if a < b {
		return a
	}
	return b
}

//const MaxInt32 = 1999999999
// 120. Triangle
func minimumTotal_dp(triangle [][]int, level int, pos int, memo [][]int) int {
	if level == len(triangle)-1 {
		return triangle[level][pos]
	}
	if memo[level][pos] == math.MaxInt32 {
		var left = minimumTotal_dp(triangle, level+1, pos, memo)
		var right = minimumTotal_dp(triangle, level+1, pos+1, memo)
		memo[level][pos] = min(left, right) + triangle[level][pos]
	}
	return memo[level][pos]
}

func minimumTotal(triangle [][]int) int {
	var memo = create2Darray(len(triangle), len(triangle), math.MaxInt32)
	return minimumTotal_dp(triangle, 0, 0, memo)
}
func main() {
	var triangle = make([][]int, 2)
	triangle[0] = make([]int, 1)
	triangle[1] = make([]int, 2)
	triangle[0][0] = 1
	triangle[1][0] = 2
	triangle[1][1] = 3
	fmt.Println(minimumTotal(triangle))
}
