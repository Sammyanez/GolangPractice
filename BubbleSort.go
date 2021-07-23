package main

func BubbleSort(array []int) []int {
	swap := true
	for swap {
		swap = false
		for i := range array {
			if i == len(array)-1 {
				break
			}
			if array[i] > array[i+1] {
				b := array[i]
				array[i] = array[i+1]
				array[i+1] = b
				swap = true
			} else {
				continue
			}
		}
	}
	return array
}
