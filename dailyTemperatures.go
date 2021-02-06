func dailyTemperatures(T []int) []int {
    count := 0
    is := false
    res := []int{}
    for pos, x := range T{
        for _, y := range T[pos+1:]{
            if x < y {
                res = append(res, count+1)
                count  = 0
                is = true
                break
            }else {
                count++
                is =false
            }
        }
        if !is {
            res = append(res, 0)
            count = 0
        }
    }
    for x := len(res); x < len(T); x++{
        res = append(res, 0)
    }
    return res
}
