package main

import "sort"

func NonConstructibleChange(coins []int) int {

	currentChange := 1

	sort.Ints(coins)

	for i := range coins{
		if currentChange >= coins[i]{
			currentChange = currentChange + coins[i]
		}
	}

 return currentChange




}

/* Solution

coins:= []int{5,7,1,1,2,3,22}
println(NonConstructibleChange(coins))

*/