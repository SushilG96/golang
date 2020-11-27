func restoreString(s string, indices []int) string {
	newStr := make([]rune, len(s))
	oldStr := []rune(s)
	for i, newIndex := range indices {
		newStr[newIndex] = oldStr[i]
	}
	return string(newStr)
}
