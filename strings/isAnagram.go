func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func isAnagram(s string, t string) bool {

	t1 := SortString(s)
	t2 := SortString(t)

    if t1 != t2 {
        return false
    }
    return true
}
