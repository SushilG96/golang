/*
Input: nums = [3,5,2,6], k = 2
Output: [2,6]

Input: nums = [2,4,3,3,5,4,9,6], k = 4
Output: [2,3,3,4]

Solution: push elements in the stack, i.e slice here. If found a element that is smaller than the last element in the stack then pop the elements, until you a found a stack of len k until the end, then return the stack.

example: 
stack:
pus pus  pop as we have 2 now   push 6
                
   |--  |                       | 
   |5   |                       |6
-- |--  |                       |--
3  |3   |2                      |2
-- |--  |--                     |--

*/
func mostCompetitive(nums []int, k int) []int {
    l := []int{}

    // Min number of element that need to be there
    m := len(nums) - k
    for _,x := range nums{
        for len(l)>0  && l[len(l)-1] > x  && m > 0 {
            m--
            l = l[:len(l)-1]
        }
        l = append(l, x)
    }
    return l[:k]
}

