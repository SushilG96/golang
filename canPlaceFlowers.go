/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.

 

Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1
Output: true
Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2
Output: false
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	if l == 1 {
		if flowerbed[0] == 0 && n >= 0 || flowerbed[0] == 1 && n == 0 {
			return true
		}
		return false
	}
	if n == 0 {
		return true
	}
	for i := 1; i < l-3; i++ {
		if flowerbed[i]+flowerbed[i+1]+flowerbed[i+2] == 0 {
			flowerbed[i+1] = 1
			n--
		}

	}
	if flowerbed[0]+flowerbed[1] == 0 {
		n--
	}
	if l > 2 {
		if flowerbed[l-1]+flowerbed[l-2] == 0 {
			n--
		}
	}
	if n == 0 || n < 0 {
		return true
	}
	return false
}
