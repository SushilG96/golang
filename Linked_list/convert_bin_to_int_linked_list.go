/*
Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.

Return the decimal value of the number in the linked list.

 

Example 1:


Input: head = [1,0,1]
Output: 5
Explanation: (101) in base 2 = (5) in base 10
Example 2:

Input: head = [0]
Output: 0


 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

import (
    "fmt"
    "strconv"
)
func getDecimalValue(head *ListNode) int {
    current := head
    s := ""
    for current != nil{
	// Convert int value to string
        s += strconv.Itoa(current.Val)
        current = current.Next
    }
    res, _ := strconv.ParseInt(s, 2, 64)
    //fmt.Println(s, res)
    return int(res)
}
