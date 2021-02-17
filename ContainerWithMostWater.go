/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.

Notice that you may not slant the container.
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
In this case, the max area of water (blue section) the container can contain is 49.
*/
func maxArea(height []int) int {
    m := 0
    for i:=0; i < len(height); i++{
        for j := i ; j < len(height); j++{
            tmp := 0
            if height[i] < height[j]{
                tmp = height[i]*(j-i)
            }else{
                tmp = height[j]*(j-i)
            }
            if tmp > m {
                m = tmp
            }
        fmt.Println(height[i],height[j],j-i,m)
        }
    }
    return m
}
/*Solution 2
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func maxArea(height []int) int {
    var res int
    var start int
    end := len(height) - 1
    for start < end {
        cur := min(height[start], height[end]) * (end - start)
        if cur > res {
            res = cur
        }
        if height[start] < height[end] {
            start++
        } else {
            end--
        }
    }
    return res
}
*/
