func searchMatrix(matrix [][]int, target int) bool {
    for _, x :=  range matrix{
        fmt.Println(x)
        if x[0] <= target && x[len(x)-1] >= target{
            for _, y := range x{
                fmt.Println(y)
                if y == target{
                    return true
                }
            }
        }
    }
    return false
}
