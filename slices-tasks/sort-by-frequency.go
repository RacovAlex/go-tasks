package slices_tasks

import "sort"

// SortByFrequency sort array by frequency of numbers.
func SortByFrequency(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	// Создаём map, чтобы подсчитать частоты чисел
	freq := make(map[int]int, 9)
	for _, v := range slice {
		freq[v]++
	}

	// Преобразуем карту в слайс структур с числом и его частотой.
	type pair struct {
		value int
		freq  int
	}
	pairs := make([]pair, 0, len(freq))
	for k, v := range freq {
		pairs = append(pairs, pair{k, v})
	}
	// Сортируем этот слайс по частоте в порядке убывания.
	// Для чисел с одинаковой частотой можно отсортировать их по возрастанию.
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].freq == pairs[j].freq {
			return pairs[i].value < pairs[j].value
		}
		return pairs[i].freq < pairs[j].freq
	})

	// Результирующий массив, отсортированный по частоте.
	out := make([]int, 0, len(slice))
	for _, v := range pairs {
		for i := 0; i < v.freq; i++ {
			out = append(out, v.value)
		}
	}
	return out
}
