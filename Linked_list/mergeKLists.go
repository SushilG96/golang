/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 
 Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
 
 */
func mergeKLists(lists []*ListNode) *ListNode {
    num := []int{}
    var res *ListNode; var head *ListNode
    for _, x := range lists{
        for x != nil {
            num = append(num, x.Val)
            x = x.Next
        }
    }
    sort.Ints(num)
    head = nil
    for n := len(num)-1; n >= 0; n-- {
        res = &ListNode{num[n], head}
        head = res
    }
    return res
}
