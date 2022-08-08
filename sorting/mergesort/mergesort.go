package mergesort

import (
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](arr []T) []T {
	return mergeSort(0, len(arr), arr)
}

func mergeSort[T constraints.Ordered](left, right int, arr []T) []T {
	if len(arr[left:right]) <= 1 {
		return arr[left:right]
	}

	mid := (left + right) / 2

	sortedLeft := mergeSort(left, mid, arr)
	sortedRight := mergeSort(mid, right, arr)

	sortedArray := make([]T, 0, len(sortedLeft)+len(sortedRight))

	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(sortedLeft) && rightIndex < len(sortedRight) {
		var value T

		if sortedLeft[leftIndex] < sortedRight[rightIndex] {
			value = sortedLeft[leftIndex]
			leftIndex++
		} else {
			value = sortedRight[rightIndex]
			rightIndex++
		}

		sortedArray = append(sortedArray, value)
	}

	for leftIndex < len(sortedLeft) {
		sortedArray = append(sortedArray, sortedLeft[leftIndex])
		leftIndex++
	}

	for rightIndex < len(sortedRight) {
		sortedArray = append(sortedArray, sortedRight[rightIndex])
		rightIndex++
	}

	return sortedArray
}
