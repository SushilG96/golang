/*
Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.

Example 1:

Input: 2
Output: [0,1,1]
Example 2:

Input: 5
Output: [0,1,1,2,1,2]
*/
func countBits(num int) []int {
	temp := []int{}

	for k := 0; k <= num; k++ {
		rem := 0
		count := 0
		i := k
		for i != 0 {
			rem = i % 2
			i /= 2
			if rem == 1 {
				count++
			}
		}
		temp = append(temp, count)
	}
	return temp
}

/*
func countBits(num int) []int {
    result := make([]int, num+1)
    result[0] = 0
    
    for i := 1; i <= num; i++ {
        if i & 1 == 1 {
            result[i] = result[i-1] + 1
        } else {
            result[i] = result[i >> 1]
        }
    }
    
    return result
}
*/
func countBitss(num int)[]int{
    fm
}
