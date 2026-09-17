func maxProfit(prices []int) int {
	l, r := 0, 1
	profit := 0
	for r < len(prices){
		if prices[l] < prices[r] {
			diff := prices[r] - prices[l]
			profit = max(profit, diff)
		} else {
			l = r
		}
		r++
	}
	return profit
}
