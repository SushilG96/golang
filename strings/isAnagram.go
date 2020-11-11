/*
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
*/
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func isAnagram(s string, t string) bool {

	t1 := SortString(s)
	t2 := SortString(t)

	if len(t1) != len(t2) {
		return false
	} else {
		for i := 0; i < len(t1); i++ {
			if t1[i] != t2[i] {
				return false
			}
		}
		return true
	}
}
