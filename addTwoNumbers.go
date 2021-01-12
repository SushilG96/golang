/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

 

Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    cur := &ListNode{}
    head := cur
    
    sum, co := 0, 0
    for co != 0 || l1 != nil || l2 != nil {
        cur.Next = &ListNode{}
        cur = cur.Next
        
        sum = co
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        
        co = sum / 10
        cur.Val = sum % 10
    }
    return head.Next
}
