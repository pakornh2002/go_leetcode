package replaceWords

import "strings"

func ReplaceWords(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")
	var builder strings.Builder

	for _, word := range words {
		replacement := word
		for _, prefix := range dictionary {
			if strings.HasPrefix(word, prefix) && len(prefix) < len(replacement) {
				replacement = prefix
			}
		}
		if builder.Len() > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(replacement)
	}

	return builder.String()
}
