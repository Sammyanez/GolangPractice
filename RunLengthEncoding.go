package main

import (
	"strconv"
	"unicode"
)

func RunLengthEncoding(str string) string {
	count := 0
	string1 := ""
	var currentChar rune = rune(str[0])
	for i, r := range str {
		if string(currentChar) == string(str[i]) && ((!unicode.IsUpper(currentChar) || unicode.IsLetter(r)) || (!unicode.IsLower(currentChar) && unicode.IsLower(r))) {
			count++
			if i == len(str)-1 {
				if count > 9 {
					for count > 9 {
						string1 += "9" + string(currentChar)
						count -= 9
					}
					string1 += strconv.Itoa(count) + string(currentChar)
				} else {
					string1 += strconv.Itoa(count) + string(currentChar)

				}
			}
		} else {
			if count > 9 {
				for count > 9 {
					string1 += "9" + string(currentChar)
					count -= 9
				}
				string1 += strconv.Itoa(count) + string(currentChar)
			} else {
				string1 += strconv.Itoa(count) + string(currentChar)

			}
			currentChar = rune(str[i])
			count = 1
			if i == len(str)-1 {
				if count > 9 {
					for count > 9 {
						string1 += "9" + string(currentChar)
						count -= 9
					}
					string1 += strconv.Itoa(count) + string(currentChar)
				} else {
					string1 += strconv.Itoa(count) + string(currentChar)

				}
			}

		}
	}
	return string1
}
