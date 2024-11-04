package interfaces

import (
	"fmt"
)

type Numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// FindMax принимает слайс любого типа (числа) и возвращает максимальный элемент слайса.
func FindMax[T Numbers](slice []T) (T, error) {
	if len(slice) == 0 {
		return *new(T), fmt.Errorf("slice is empty")
	}

	maximum := slice[0]
	for _, value := range slice {
		if value > maximum {
			maximum = value
		}
	}
	return maximum, nil
}

// Filter принимает слайс элементов любого типа и функцию-предикат,
// возвращает новый слайс, содержащий только те элементы, для которых предикат возвращает true.
func Filter[T comparable](s []T, fn func(T) bool) []T {
	result := make([]T, 0, len(s))
	for _, value := range s {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}

// Reduce принимает слайс элементов любого типа, функцию, которая последовательно применяет некоторое действие к элементам,
// и начальное значение. Функция возвращает результат выполнения этой операции над всеми элементами слайса.
func Reduce[T any](s []T, start T, fn func(a, b T) T) T {
	result := start
	for _, value := range s {
		result = fn(result, value)
	}
	return result
}
