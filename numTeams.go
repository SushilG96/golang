/*
There are n soldiers standing in a line. Each soldier is assigned a unique rating value.

You have to form a team of 3 soldiers amongst them under the following rules:

Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rating[k]).
A team is valid if:  (rating[i] < rating[j] < rating[k]) or (rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n).
Return the number of teams you can form given the conditions. (soldiers can be part of multiple teams).

 

Example 1:

Input: rating = [2,5,3,4,1]
Output: 3
Explanation: We can form three teams given the conditions. (2,3,4), (5,4,1), (5,3,1). 
*/
func numTeams(rating []int) int {
    le := len(rating)
    count := 0
    for i := 0; i < le; i++ {
		for j := i + 1; j < le; j++ {
			for k := j + 1; k < le; k++ {
				if (rating[i] > rating[j] && rating[j] > rating[k]) || (rating[i] < rating[j] && rating[j] < rating[k]) {
					count++
                    //fmt.Println(nums[i], nums[j], nums[k])
				}
			}
		}
	}
    return count
}
/*Sol 2

func numTeams(rating []int) int {
	result := 0
	if len(rating) == 0 {
		return result
	}

	n := len(rating)
	up := make([]int, n)
	down := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if rating[i] < rating[j] {
				up[i] += 1
				result += up[j]
			} else {
				down[i] += 1
				result += down[j]
			}
		}
	}
	return result
}
*/
