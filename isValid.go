func isValid(s string) bool {
	// last open stack
	var los []rune

	opposite := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			los = append(los, c)
		} else {
			if len(los) > 0 && c == opposite[los[len(los)-1]] {
				los = los[:len(los)-1]
			} else {
				return false
			}
		}
	}

	return len(los) == 0
}
