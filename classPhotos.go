package main

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	iterator := 0
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	if redShirtHeights[iterator] == blueShirtHeights[iterator] {
		return false
	} else if redShirtHeights[iterator] > blueShirtHeights[iterator] {
		for i, x := range redShirtHeights {
			if x <= blueShirtHeights[i] {
				return false
			}
		}
	} else {
		for i, x := range redShirtHeights {
			if x >= blueShirtHeights[i] {
				return false
			}
		}
	}
	return true
}
