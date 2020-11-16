func topKFrequent(nums []int, k int) []int {

	if len(nums) == k {
		return nums
	}
	tmp := make(map[int]int)
	res := []int{}
	for _, i := range nums {
		tmp[i]++
	}
	value := []int{}
	for _, k := range tmp {
		value = append(value, k)
	}
	sort.Ints(value)

	// reverse slice
	for i, j := 0, len(value)-1; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}

	for i := k - 1; i >= 0; i-- {
		for j, _ := range tmp {
			if tmp[j] == value[i] {
				res = append(res, j)
			}
		}

	}
	// Make slice distinct
	u := make([]int, 0, k)
	m := make(map[int]bool)

	for _, val := range res {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u

}
