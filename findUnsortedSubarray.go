/*
Given an integer array nums, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order.

Return the shortest such subarray and output its length.

 

Example 1:

Input: nums = [2,6,4,8,10,9,15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
*/
func findUnsortedSubarray(nums []int) int {
	l := len(nums)
	tmp := make([]int, l)

	copy(tmp, nums)
	sort.Ints(tmp)

	start := 0
	end := 0
	for i := 0; i < l; i++ {
		if nums[i] != tmp[i] {
			start = i + 1
			break
		}
	}
	for j := l - 1; j > 0; j-- {
		if nums[j] != tmp[j] {
			end = j + 1
			break

		}

	}
	if start == 0 && end == 0 {
		return 0
	}
	return end - start + 1
}

