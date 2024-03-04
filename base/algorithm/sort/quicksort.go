package sort

// 2 1 5 3 4

// 4 1 5 3 2

// 3 1 5 4 2

// 1 2 3 4 5

func QuickSort[T int | string](items []T) {
	quickSort(items, 0, len(items)-1)
}

func quickSort[T int | string](items []T, left, right int) {
	if left >= right {
		return
	}
	pivot := partition2(items, left, right)
	quickSort(items, left, pivot-1)
	quickSort(items, pivot+1, right)
}

func partition[T int | string](items []T, left, right int) int {
	i, j := left, right

	for i < j {
		for i < j && items[j] >= items[left] {
			j-- // 从右向左找首个小于基准数的元素
		}
		for i < j && items[i] <= items[left] {
			i++ // 从左向右找首个大于基准数的元素
		}
		// 元素交换
		items[i], items[j] = items[j], items[i]
	}
	// 将基准数交换至两子数组的分界线
	items[i], items[left] = items[left], items[i]
	return i // 返回基准数的索引
}

func partition2[T int | string](items []T, left, right int) int {
	i, j := left, right

	for i < j {
		for i < j {
			if items[j] < items[left] {
				break
			} else {
				j--
			}
		}

		for i < j {
			if items[i] > items[left] {
				break
			} else {
				i++
			}
		}
		items[i], items[j] = items[j], items[i]
	}
	items[i], items[left] = items[left], items[i]

	return i
}
