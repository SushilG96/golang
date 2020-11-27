/*
Given the array arr of positive integers and the array queries where queries[i] = [Li, Ri], for each query i compute the XOR of elements from Li to Ri (that is, arr[Li] xor arr[Li+1] xor ... xor arr[Ri] ). Return an array containing the result for the given queries.
 

Example 1:

Input: arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
Output: [2,7,14,8] 
Explanation: 
The binary representation of the elements in the array are:
1 = 0001 
3 = 0011 
4 = 0100 
8 = 1000 
The XOR values for queries are:
[0,1] = 1 xor 3 = 2 
[1,2] = 3 xor 4 = 7 
[0,3] = 1 xor 3 xor 4 xor 8 = 14 
[3,3] = 8
*/
func xorQueries(arr []int, queries [][]int) []int {
	final := make([]int, len(queries))
	for pos, x := range queries {
		count := 0
		if x[0] != 0 || x[1] != 0 {
			for i := x[0]; i <= x[1]; i++ {
				count ^= arr[i]
			}
		} else {
			count = arr[0]
		}
		final[pos] = count
	}
	return final
}
/*
func xorQueries(arr []int, queries [][]int) []int {
    
    for i:=1;i<len(arr);i++{
        arr[i] = arr[i]^arr[i-1]
    }
    res:=[]int{}
    for i:=0;i<len(queries);i++ {
        if queries[i][0] == 0 {
            res = append(res, arr[queries[i][1]])
        } else {
            res = append(res, arr[queries[i][1]] ^ arr[queries[i][0]-1])
        }
    }
    
    return res
    
}
*/