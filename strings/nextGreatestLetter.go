/*
Given a list of sorted characters letters containing only lowercase letters, and given a target letter target, find the smallest element in the list that is larger than the given target.

Letters also wrap around. For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.

Examples:
Input:
letters = ["c", "f", "j"]
target = "a"
Output: "c"

Input:
letters = ["c", "f", "j"]
target = "j"
Output: "c"

Input:
letters = ["c", "f", "j"]
target = "k"
Output: "c"


*/
func nextGreatestLetter(letters []byte, target byte) byte {
    for _, letter := range letters{
        // Return if found the smallest greater than target
        if letter > target{
             return letter
        }
    }

    // If not then return the first number
   return letters[0]
}
