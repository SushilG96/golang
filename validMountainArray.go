/*
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < A[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
*/
func validMountainArray(arr []int) bool {
	max := arr[0]; pos := 0; top := len(arr)
	for i := 1; i < top; i++ {
		if arr[i] > max {
			max = arr[i]
			pos = i
		}
	}
	if pos == 0 || pos == top-1 {
		return false
	}
	for i := 1; i < pos; i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	for j := pos; j < top-1; j++ {
		if arr[j+1] >= arr[j] {
			return false
		}
	}
	return true
}
