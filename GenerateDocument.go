package main

func GenerateDocument(characters string, document string) bool {

	var b = make([]byte, len(characters))
	copy(b, characters)

	var c = make([]byte, len(document))
	copy(c, document)

	for i := range b {
		for j := range c {
			if c[j] == b[i] {
				c = append(c[:j], c[j+1:]...)
				break
			}
		}
	}
	if len(c) == 0 {
		return true
	}

	return false
}
