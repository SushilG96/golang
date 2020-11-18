/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

 

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

*/
func merge(intervals [][]int) [][]int {

	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		pos := len(res) - 1
		if res[pos][0] <= intervals[i][0] && intervals[i][0] <= res[pos][1] {
			if intervals[i][1] > res[pos][1] {
				res[pos][1] = intervals[i][1]
			}

		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
