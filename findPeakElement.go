package main

import ("fmt")

func main(){
    n := []int{1,2,1,3,5,6,4}
    fmt.Println(findPeakElement(n))
}
/*
A peak element is an element that is strictly greater than its neighbors.

Given an integer array nums, find a peak element, and return its index.
If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞.
Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
*/
func findPeakElement(nums []int) int {
    m := -1; i := 0
    for pos, x := range nums{
        if x > m{
            m = x
            i = pos
       }
    }
    return i
}
