/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	tmpList := []int{}

	var res *ListNode

	for current != nil {
		tmpList = append(tmpList, current.Val)
		current = current.Next
	}

	m := make(map[int]int)

	for _, v := range tmpList {
		m[v]++
	}

    finalList := []int{}
	for x, y := range m {
		if y == 1 {
			finalList = append(finalList, x)

		}
	}
	sort.Ints(finalList)
	head = nil
	for x := len(finalList) - 1; x >= 0; x-- {
		res = &ListNode{finalList[x], nil}
		res.Next = head
		head = res
	}

	return res
}

