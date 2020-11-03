/*

Given a string s, find the length of the longest substring without repeating characters.

 

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
Example 4:

Input: s = ""
Output: 0
*/
func lengthOfLongestSubstring(s string) int {
	max := 0
	store := []byte{}
	for _, i := range s {
		if bytes.ContainsAny(store, string(i)) == false {
			store = append(store, byte(i))
			if len(store) > max {
				max = len(store)
			}
		} else {
			pos := 0
			if len(store) > 1 {
				for p, j := range store {
					if j == byte(i) {
						pos = p
					}
				}
			}
			store = append(store[pos+1:], byte(i))

		}
	}
	if len(store) > max {
		return len(store)
	}
	return max
}
