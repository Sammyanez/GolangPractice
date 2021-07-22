package main

import "sort"

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	if fastest {
		sort.Sort(sort.Reverse(sort.IntSlice(blueShirtSpeeds)))
	} else {
		sort.Ints(blueShirtSpeeds)
	}
	realMax := 0

	for i, x := range redShirtSpeeds {
		if x > blueShirtSpeeds[i] {
			realMax += x

		} else {
			realMax += blueShirtSpeeds[i]

		}
	}

	return realMax

}
