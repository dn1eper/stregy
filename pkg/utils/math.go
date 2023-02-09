package utils

func Max(nums ...float64) float64 {
	max := 0.0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func Min(nums ...float64) float64 {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}

	return min
}
