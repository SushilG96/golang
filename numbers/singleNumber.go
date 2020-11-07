func singleNumber(nums []int) int {
	x := 0
	for _, i := range nums {
		x ^= i
	}
	return x

}
