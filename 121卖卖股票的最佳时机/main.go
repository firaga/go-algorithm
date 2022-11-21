package _21卖卖股票的最佳时机

//穷举
func maxProfit0(prices []int) int {
	var min = 0
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			if prices[j]-prices[i] > min {
				min = prices[j] - prices[i]
			}
		}
	}
	return min
}

//同时计算最小值和最大收益
func maxProfit1(prices []int) int {
	var min = prices[0]
	var max = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}

//动态规划
//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	min := prices[0]
	i := 1
	for ; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if dp[i-1] > prices[i]-min {
			dp[i] = dp[i-1]
		} else {
			dp[i] = prices[i] - min
		}
	}
	return dp[i-1]
}
