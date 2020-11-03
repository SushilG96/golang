/*Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]*/
func findDisappearedNumbers(nums []int) []int {
    length := len(nums)
    if length == 0 { return nil }
    res := make([]int, length)
    for _, v := range nums {
        res[v-1] = v
    }
    j := 0
    for i := 0; i < length; i++ {
        if res[i] == 0 {
            res[j] = i + 1
            j++
        }
    }
    return res[:j]
}
