/*
Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example 1:

Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL
*/
func oddEvenList(head *ListNode) *ListNode {
  if head == nil {
    return nil
  } else if head.Next == nil {
    return head
  }

  oddsHead := head
  evensHead := head.Next  
  for curr, temp := head, head.Next; temp != nil; curr, temp = temp, temp.Next {
    curr.Next = temp.Next
  }
  
  oddsTail := oddsHead
  for ;oddsTail.Next != nil; oddsTail = oddsTail.Next {}
  oddsTail.Next = evensHead
  
  return oddsHead
}
