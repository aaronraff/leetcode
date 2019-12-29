/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // Need a dummy to handle edge cases, like the
    // nth node being at the beginning of the list
    dummy := &ListNode { Next: head, Val: -1 }
    left := dummy
    right := dummy
    
    // Create a gap of n between the two pointers
    for i := 0; i < n+1; i++ {
        right = right.Next
    }
    
    // Once right is at the end of the list, left will be
    // right before the node we want to remove
    for right != nil {
        right = right.Next
        left = left.Next
    }
        
    // Skip over the nth node from the end
    left.Next = left.Next.Next  
    return dummy.Next
}
