package main

func FirstNonRepeatingCharacter(str string) int {
	var s string
	for i := 0; i < len(str) - 1; i++ {
		for j := i + 1; j <= len(str) ; j++ {
			if j == len(str) {
				return i
			}else if str[i] == str[j] {
				s = s + string(str[j])
				break
					}
			for k := 0; k < len(s); k++ {
				if str[i] == s[k] {
					i++
					j = i
					}

			}
				}
			}
	if len(str) == 1{
		return 0
	}
	return -1
}
