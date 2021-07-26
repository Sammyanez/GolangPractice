package main

func SelectionSort(array []int) []int {

	for i := range array {
		smallest := array[i]
		smallestIndx := i
		for j := i + 1; j < len(array); j++ {
			if smallest > array[j] {
				smallest = array[j]
				smallestIndx = j
			}
		}
		temp := array[i]
		array[i] = array[smallestIndx]
		array[smallestIndx] = temp

	}
	return array
}
