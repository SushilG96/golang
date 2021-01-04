// Sol 1
func findMaxAverage(nums []int, k int) float64 {

	m := -999999.0
	for i := 0; i <= len(nums)-k; i++ {
		avg := 0.0
		for _, j := range nums[i : i+k] {
			avg += float64(j)
		}
		if avg/float64(k) > m {
			m = avg / float64(k)
		}
	}
	return m
}

// Sol 2
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}
	res := sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		res = max(sum, res)
	}
	return float64(res) / float64(k)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
