package driver

import "math"

func max(nums []int) int {
	maxVal := math.MinInt32
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}
