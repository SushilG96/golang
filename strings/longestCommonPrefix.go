/*Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

 

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	var prefix string
	str := strs[0]
	for i := 0; i < len(str); i++ {
		prefix = string(str[:i+1])
		for _, s := range strs {
			if !strings.HasPrefix(s, prefix) {
				return string(prefix[:len(prefix)-1])
			}
		}
	}
	return prefix
}

