func kidsWithCandies(candies []int, extraCandies int) []bool {
	m := candies[0]
	tmp := make([]bool, len(candies))
	for i := 1; i < len(candies); i++ {
		if m < candies[i] {
			m = candies[i]
		}
	}

	for pos, c := range candies {
		if c+extraCandies >= m {
			tmp[pos] = true
		} else {
			tmp[pos] = false
		}
	}
	return tmp
}
