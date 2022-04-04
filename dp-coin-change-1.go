package main

import (
	"fmt"
)

/**
* Find total number of combinations to make the given sum
* ex
* coin denominations = [1,3,5]
* sum = 10
 */
var cache = [][]int{}

func main() {
	coins := []int{1, 3, 5}
	const sum int = 10

	initCache(sum, len(coins))
	fmt.Println(totalCombinations(coins, sum, len(coins)-1))
}

func initCache(sum int, coins int) {
	for i := 0; i < sum+1; i++ {
		row := []int{}
		for j := 0; j < coins; j++ {
			row = append(row, -1)
		}
		cache = append(cache, row)
	}
}

func totalCombinations(coins []int, sum int, i int) int {
	if i < 0 || sum < 0 {
		return 0
	}
	if cache[sum][i] != -1 {
		return cache[sum][i]
	}
	if sum == 0 {
		return 1
	}
	cache[sum][i] = totalCombinations(coins, sum-coins[i], i) + totalCombinations(coins, sum, i-1)
	return cache[sum][i]
}
