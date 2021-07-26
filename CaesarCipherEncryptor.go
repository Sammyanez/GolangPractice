package main

/*
func CaesarCipherEncryptor(str string, key int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var myNums []int
	var myString string
	for i := range str{
		for j:= range alphabet{
			if str[i] == alphabet[j] {
				myNums = append(myNums, j)
				break
			}
		}
	}
	for x := range myNums {
		myNums[x] += key
		for myNums[x] > 25 {
			myNums[x] -= 26
		}
		myString += toCharStr(myNums[x])
	}

	return myString
}

func toCharStr(i int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	return string(alphabet[i])
}
*/

func CaesarCipherEncryptor(str string, key int) string {
	shift, offset := rune(key%26), rune(26)
	runes := []rune(str)
	for i, char := range runes {
		if char >= 'a' && char+shift <= 'z' {
			char += shift
		} else {
			char += shift - offset
		}
		runes[i] = char
	}
	return string(runes)
}
