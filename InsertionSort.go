package main

func InsertionSort(array []int) []int {
	length := len(array)
	for i := range array {
		if i != length-1 {
			if array[i] > array[i+1] {
				c := array[i]
				array[i] = array[i+1]
				array[i+1] = c
			}
		}
		for j := i; j >= 0; j-- {
			if j-1 == -1 {
				break
			}
			if array[j] < array[j-1] {
				b := array[j]
				array[j] = array[j-1]
				array[j-1] = b
			} else {
				continue
			}
		}

	}
	return array
}
