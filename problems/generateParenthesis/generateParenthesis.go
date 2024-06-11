package generateparenthesis

func GenerateParenthesis(n int) []string {
	var result []string

	var backtrack func(curr string, open int, close int)
	backtrack = func(curr string, open int, close int) {
		if len(curr) == n*2 {
			result = append(result, curr)
			return
		}
		if open < n {
			backtrack(curr+"(", open+1, close)
		}
		if close < open {
			backtrack(curr+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return result
}
