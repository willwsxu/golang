// 474. Ones and Zeroes
func max(a int, b int) int {
    if a>b { return a}
    return b
}

func findMaxForm(strs []string, m int, n int) int {
    memo := make([][]int, m+1)
    for x:= range memo {
        memo[x]=make([]int, n+1)
    }
    for i:=0; i<len(strs); i++ {
        var ones, zero int =0,0
        for j:=0; j<len(strs[i]); j++ {
            if strs[i][j]=='1' { ones++ } else {zero++}
        }
        for j:=m; j>=zero; j-- {
            for k:=n; k>=ones; k-- {
                memo[j][k] = max(memo[j][k], 1+memo[j-zero][k-ones])
            }
        }
    }
    return memo[m][n]
}

// 377. Combination Sum IV
var dp []int
func combinationSum4Dp(nums []int, target int) int {
    if dp[target]>=0 { return dp[target]}
    var total int = 0
    for i:=0; i<len(nums); i++ {
        if target>= nums[i] {
            total += combinationSum4Dp(nums, target-nums[i])
        }
    }
    dp[target]=total
    return total
}

func combinationSum4(nums []int, target int) int {
    dp=make([]int, target+1)
    for i:=0; i<=target; i++ {
        dp[i]=-1
    }
    dp[0]=1
    return combinationSum4Dp(nums, target)
}