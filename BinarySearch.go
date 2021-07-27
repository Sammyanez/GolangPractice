package main

func BinarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1

	for left < right {
		middle := (left + right) / 2
		if array[middle] < target {
			left = middle + 1
		} else if array[middle] > target {
			right = middle - 1
		} else {
			return middle
		}
		if left == right && array[left] == target {
			return left
		}

	}

	return -1
}
