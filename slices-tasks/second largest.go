package slices_tasks

// Найти второе по максимальности число в массиве чисел

func SecondLargest(nums []int) (int, bool) {
	if len(nums) < 2 {
		return 0, false
	}
	maximum, secondMax := nums[0], nums[0]
	for _, num := range nums {
		if num > maximum {
			secondMax = maximum
			maximum = num
		} else if num > secondMax && num != maximum {
			secondMax = num
		}
	}

	if maximum == secondMax {
		return 0, false
	}
	return secondMax, true
}
