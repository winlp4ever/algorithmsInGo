package main

func BuyAndSellStocks(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	b := nums[0]
	profit := 0
	for i, u := range nums {
		if u < b {
			b = u
		} else if (i < n-1 && u > nums[i+1]) || i == n-1 {
			profit += u - b 
			if (i < n-1) {
				b = nums[i+1]
			}
		}
	}
	return profit
}