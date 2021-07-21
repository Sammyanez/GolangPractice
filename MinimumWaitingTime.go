package main

import "sort"

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)

	minimum := 0
	MinimumWaitingTimeCalculator(queries,&minimum)

	return minimum
}

func MinimumWaitingTimeCalculator(queries []int, i *int) {
	for x := 0; x < len(queries) - 1; x++{
		*i += queries[x] * (len(queries) -x -1)
	}

}
