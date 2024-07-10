package reverseword

func ReverseWord(word string) string {
	result := ""
	start := 0

	for i := 0; i <= len(word); i++ {
		if i == len(word) || word[i] == ' ' {
			for j := i - 1; j >= start; j-- {
				result += string(word[j])
			}
			if i < len(word) {
				result += " "
			}
			start = i + 1
		}
	}

	return result
}
