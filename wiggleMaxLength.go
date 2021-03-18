package main

import (
    "fmt"
)
func wiggleMaxLength(nums []int) int {
    res := []int {}
    for x := 1; x < len(nums); x++{
        cal := nums[x-1]-nums[x]
        if cal != 0{res = append(res, cal)}
    }
    if len(res) == 0{return 1}

    r := 2
    for x := 1; x < len(res); x++{
        if res[x-1]>0 && res[x]<0  || res[x-1]<0 && res[x]>0{r++}
    }
    fmt.Println(res, r)
    return r
}
func main(){
    nums := []int {1, 7, 4, 9, 2, 5}
    fmt.Println(wiggleMaxLength(nums))
}
