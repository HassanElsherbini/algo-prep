package leetcode

import "github.com/HassanElsherbini/algo-prep/common"

/* EASY

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future
to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.


Solution 1 (brute force - evaluate all combinations): Time O(N^2) | Space O(1)
- for each price in array (buy price b):
	- for each remaining prices (sell price s):
		- calculate profit: s - b
		- update max profit if new profit is larger
- return max profit
--------------------------------------------------------------------------------------

Solution 2: Time O(N) | Space O(1)
- initialize buy price to first price and max profit to 0
- for each remaining price:
	- if price is less than buy price:
		- no profit to make but better buy price, so set price as new buy price
	- else if price > buy price && profit(price-buy price) > max profit:
			- set profit as new max profit
- return max profit

*/

func BestTimeToBuySellStock(prices []int) int {
	var maxProfit int
	buyPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		sellPrice := prices[i]
		if prices[i] < buyPrice {
			buyPrice = sellPrice
		} else if sellPrice > buyPrice {
			maxProfit = common.Max(maxProfit, sellPrice-buyPrice)
		}
	}

	return maxProfit
}
