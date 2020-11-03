/*Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand. (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Return the minimum element of this array.

 

Example 1:

Input: nums = [3,4,5,1,2]
Output: 1
Example 2:

Input: nums = [4,5,6,7,0,1,2]
Output: 0
Example 3:

Input: nums = [1]
Output: 1*/
func findMin(nums []int) int {
    m := nums[0]
    for _,i:=range nums[1:]{
        if i < m{
            m = i
        }
    }
    return m
}
