package sort

import "cmp"

func QuickSort[T cmp.Ordered](items []T) {
	quickSort(items, 0, len(items)-1)
}

func quickSort[T cmp.Ordered](items []T, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(items, left, right)
	quickSort(items, left, pivot-1)
	quickSort(items, pivot+1, right)
}

func partition[T cmp.Ordered](items []T, left, right int) int {
	for left < right {
		for left < right && items[right] >= items[left] {
			right--
		}
		items[left], items[right] = items[right], items[left]

		for left < right && items[left] <= items[right] {
			left++
		}
		items[left], items[right] = items[right], items[left]
	}
	return left
}
