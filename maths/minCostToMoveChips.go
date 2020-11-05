/*We have n chips, where the position of the ith chip is position[i].

We need to move all the chips to the same position. In one step, we can change the position of the ith chip from position[i] to:

position[i] + 2 or position[i] - 2 with cost = 0.
position[i] + 1 or position[i] - 1 with cost = 1.
Return the minimum cost needed to move all the chips to the same position.

 

Example 1:


Input: position = [1,2,3]
Output: 1
Explanation: First step: Move the chip at position 3 to position 1 with cost = 0.
Second step: Move the chip at position 2 to position 1 with cost = 1.
Total cost is 1.
*/
func minCostToMoveChips(position []int) int {
	odd := 0
	even := 0

	for _, i := range position {
		if i%2 == 0 {
			even++
		} else {
			odd++
		}

	}
	if even == len(position) || odd == len(position) {
		return 0
	} else {
		if even > odd {
			return odd
		}
	return even
	}
}
