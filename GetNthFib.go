package main

func GetNthFib(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 0
	} else {
		return GetNthFib(n-1) + GetNthFib(n-2)
	}

}
