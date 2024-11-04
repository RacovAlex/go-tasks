package slices_tasks

import (
	"fmt"
)

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//Note that you must do this in-place without making a copy of the array.

func MoveZeroes(nums []int) {
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}

// Для задачи перемещения всех нулей в конец массива без изменения порядка других элементов
// и при этом в режиме in-place обычно применяется алгоритм с использованием двух указателей.
// Один указатель (nonZeroIndex) отслеживает позицию, куда необходимо переместить следующий ненулевой элемент.
// Второй указатель (i) проходит по всему массиву и ищет ненулевые элементы.
