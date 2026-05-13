package best_time_stock

import "math"

// MaxProfit returns the maximum profit from one buy-sell transaction.
// You must buy before you sell. Return 0 if no profit is possible.
// Input: [7,1,5,3,6,4]  →  5  (buy at 1, sell at 6)
func MaxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0 

		for _, price := range prices {
			if price < minPrice {
				minPrice = price
			} else if price - minPrice > maxProfit {
				maxProfit = price - minPrice
			}
		}
	return maxProfit
}
