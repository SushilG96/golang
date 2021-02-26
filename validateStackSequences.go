/*
Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
Output: true
Explanation: We might do the following sequence:
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
*/
func validateStackSequences(pushed []int, popped []int) bool {
    stack := []int{}
    pos := 0
    for _, x := range pushed{
        stack = append(stack, x)
        for len(stack)> 0 && stack[len(stack)-1] == popped[pos]{
            pos++
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}
