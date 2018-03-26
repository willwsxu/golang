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

func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}

func min(a int, b int) int { // math.Min is float type
	return Ternary(a < b, a, b).(int)
}

//const MaxInt32 = 1999999999
// 120. Triangleternary
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

func testTriangle() {
	var triangle = make([][]int, 2)
	triangle[0] = make([]int, 1)
	triangle[1] = make([]int, 2)
	triangle[0][0] = 1
	triangle[1][0] = 2
	triangle[1][1] = 3
	fmt.Println(minimumTotal(triangle))
}

// 416. Partition Equal Subset Sum
func canPartition(nums []int) bool {
	var sum = 0
	for _, n := range nums {
		sum += n
	}
	if sum&1 == 1 {
		return false
	} //odd
	sum /= 2
	var dp = make([]bool, sum+1)
	dp[0] = true
	//fmt.Printf("%v\n", dp)
	for _, n := range nums {
		if sum >= n {
			dp[sum] = dp[sum] || dp[sum-n]
			if dp[sum] {
				return true
			}
		}
		for j := sum - 1; j > 0; j-- {
			if j >= n {
				dp[j] = dp[j] || dp[j-n]
			}
		}
	}
	return dp[sum]
}

func testEqualSumPartition() {
	var nums = [...]int{1, 5, 11, 5}          //array
	fmt.Printf("%t\n", canPartition(nums[:])) // convert to slice

	var nums2 = [...]int{1, 2, 3, 5}
	fmt.Printf("%t\n", canPartition(nums2[:]) == false)

	var nums3 = [...]int{3, 3, 3, 4, 5}
	fmt.Printf("%t\n", canPartition(nums3[:]))
}

func dfsKPartition(nums []int, k int, chosen []bool, subsetSum int, target int, setIdx int, numIdx int) bool {
	if k == 1 && target != 0 || k == 0 {
		return true
	}
	if subsetSum == target && setIdx > 0 { // setIdx>0 is used to support target=0
		return dfsKPartition(nums, k-1, chosen, 0, target, 0, 0)
	}
	for i := numIdx; i < len(nums); i++ {
		if chosen[i] || subsetSum+nums[i] > target {
			continue
		}
		chosen[i] = true
		if dfsKPartition(nums, k, chosen, subsetSum+nums[i], target, setIdx+1, i+1) {
			return true
		}
		chosen[i] = false
	}
	return false
}
func canPartitionKSubsets(nums []int, k int) bool {
	if k < 1 {
		return false
	}
	if k == 1 {
		return true
	}
	var sum = 0
	for _, n := range nums {
		sum += n
	}
	if sum%k > 0 {
		return false
	} //not divisible
	var chosen = make([]bool, len(nums))
	return dfsKPartition(nums, k, chosen, 0, sum/k, 0, 0)
}
func testKPartition() {
	var nums3 = [...]int{3, 3, 3, 4, 5}
	fmt.Printf("%t\n", canPartitionKSubsets(nums3[:], 2))
}

// a 2x1 domino shape, and an "L" tromino shape
// Given N, how many ways are there to tile a 2 x N board? Return your answer modulo 10^9 + 7
func numTilings(N int) int {
	const MOD int = 1000000007
	if N < 3 {
		return N
	}
	var dp = make([]int, N+1) // flat end
	dp[0] = 0
	dp[1] = 1 // n=1, one way
	dp[2] = 2
	var dp2 = make([]int, N+1) // tromino end
	dp2[0] = 0
	dp2[1] = 0
	dp2[2] = 2
	for n := 3; n <= N; n++ {
		dp[n] = (dp[n-1] + dp[n-2]) % MOD  // domino to domino
		dp[n] = (dp[n] + dp2[n-1]) % MOD   // tromino+tromino
		dp2[n] = (2 * dp[n-2]) % MOD       // domino + tromino
		dp2[n] = (dp2[n] + dp2[n-1]) % MOD // tromino + domino extend tromino
	}
	return dp[N]
}

func tern_op(cond bool, first, second int) int {
	if cond {
		return first
	}
	return second
}
func findPaths(m int, n int, N int, i int, j int) int {
	var dp3 = make([][][]int, 2)
	const SIZE int = 50
	for i := 0; i < 2; i++ {
		dp3[i] = make([][]int, SIZE)
		for j := 0; j < SIZE; j++ {
			dp3[i][j] = make([]int, SIZE)
			//fmt.Println("i %v j=%v, array=%v", i, j, dp3[i][j])
		}
	}
	const MOD int = 1000000007
	var prev int = 0
	var current int = 0
	for x := 1; x <= N; x++ {
		current = 1 - prev
		//fmt.Println("loop %v current=%v", x, current)
		for r := 0; r < m; r++ {
			for c := 0; c < n; c++ {
				//fmt.Println("loop %v current=%v r=%v c=%v", x, current, r, c)
				if r == 0 { // up
					dp3[current][r][c] = 1
				} else {
					dp3[current][r][c] = dp3[prev][r-1][c]
				}
				if r == m-1 {
					dp3[current][r][c] += 1
				} else {
					dp3[current][r][c] += dp3[prev][r+1][c] // down
					dp3[current][r][c] %= MOD
				}
				if c == 0 {
					dp3[current][r][c] += 1
				} else {
					dp3[current][r][c] += dp3[prev][r][c-1] // left
					dp3[current][r][c] %= MOD
				}
				if c == n-1 {
					dp3[current][r][c] += 1
				} else {
					dp3[current][r][c] += dp3[prev][r][c+1] // left
					dp3[current][r][c] %= MOD
				}
			}
		}
		prev = current // swap
	}
	return dp3[current][i][j]
}
func main() {
	fmt.Println("%v\n", findPaths(2, 2, 2, 0, 0))
}
