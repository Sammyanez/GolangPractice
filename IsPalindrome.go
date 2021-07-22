package main

func IsPalindrome(str string) bool {
	for x := 0; x < len(str); x++ {
		if str[x] != str[len(str)-1-x] {
			return false
		}
	}

	return true
}
