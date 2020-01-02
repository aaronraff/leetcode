/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{ Next: head }
    curr := dummy
    
    // The next two nodes are not nil,
    // meaning we can swap them
    for curr.Next != nil && curr.Next.Next != nil {
        left := curr.Next
        right := left.Next
        curr.Next = right
        left.Next = right.Next
        right.Next = left
        
        curr = curr.Next.Next
    }
    
    return dummy.Next
}
