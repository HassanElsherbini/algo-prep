package dp

import (
	"github.com/HassanElsherbini/algo-prep/common"
)

/* MEDIUM

SOLUTION 1 (brute force): Time O(2^N) | Space O(N)
	- if current capacity <= 0 or there is no more items to process:
		- return 0
	- if current item <= capacity:
		- calculate profit for selecting this item
	- calcualate profit for skipping this item
	- return the larger profit of the two

----------------------------------------------------------------------------
SOLUTION 2 (same as solution 1 but using memoization (dp top-down)): Time O(N * C) | Space O(N * C) C = capacity
(all weight/capacity combinations are evaluated once)
	- if current capacity <= 0 or there is no more items to process:
		- return 0
	- if a max profit for current capacity/item exists in memo:
		- return it
	- if current item <= capacity:
		- calculate profit for selecting this item
	- calcualate profit for skipping this item
	- store max of those two profits in memo for the current item/capacity
	- return such max

----------------------------------------------------------------------------
SOLUTION 3 (dp bottom-up approach): Time O(N * C) | Space O(C) C = capacity
	- create an array that will hold the running max profit for capacities 0 --> C
	- initialize that array with the max profit possible if only the first item can be selected
	- for each remaining item:
		- for each capacity starting from input C --> 0:
			- if item can be selected:
				- calculate profit for selecting the item by adding the item profit + the current max profit
					at the remaining capacity(capacity - item weight)
			- max profit at capacity = Max(the new calculated profit, current max profit at capacity)

	- return max profit at highest capcity (input C)
*/

func KnapSackTopDown(profits, weights []int, capacity int) int {
	memo := make([][]int, len(weights))
	for i := range memo {
		memo[i] = make([]int, capacity)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return maxProfit(profits, weights, capacity, 0, memo)
}

func maxProfit(profits, weights []int, capacity, index int, memo [][]int) int {
	if capacity <= 0 || index == len(weights) {
		return 0
	}

	if memo[index][capacity-1] > -1 {
		return memo[index][capacity-1]
	}

	var profitSelectingItem, profitSkippingItem int
	if weights[index] <= capacity {
		profitSelectingItem = profits[index] + maxProfit(profits, weights, capacity-weights[index], index+1, memo)
	}
	profitSkippingItem = maxProfit(profits, weights, capacity, index+1, memo)

	memo[index][capacity-1] = common.Max(profitSelectingItem, profitSkippingItem)
	return memo[index][capacity-1]
}

func KnapSackBottomUp(profits, weights []int, capacity int) int {
	maxRunningProfitAtCapacity := make([]int, capacity+1)
	for i := range maxRunningProfitAtCapacity {
		if weights[0] <= i {
			maxRunningProfitAtCapacity[i] = profits[0]
		}
	}
	for i := 1; i < len(weights); i++ {
		for c := capacity; c >= 0; c-- {
			var profitSelectingItem int
			if weights[i] <= c {
				profitSelectingItem = profits[i] + maxRunningProfitAtCapacity[c-weights[i]]
			}
			maxRunningProfitAtCapacity[c] = common.Max(profitSelectingItem, maxRunningProfitAtCapacity[c])
		}
	}

	return maxRunningProfitAtCapacity[capacity]
}
