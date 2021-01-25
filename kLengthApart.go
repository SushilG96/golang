/*
Input: nums = [1,0,0,1,0,1], k = 2
Output: false
*/
func kLengthApart(nums []int, k int) bool {
    first_1 := true; start := 0
    for pos, x := range nums{
        if x == 1 && first_1 {
            start=pos
            first_1 = false
        }else if x == 1  {
            if start >= k{
                start = 0
            }else{return false}
        }else{
            start++
        }
    }
    return true
}

