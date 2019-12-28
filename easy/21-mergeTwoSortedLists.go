/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    curr := head
        
    for l1 != nil || l2 != nil {
        if l1 == nil {
            curr.Next = l2
            l2 = l2.Next
        } else if l2 == nil {
            curr.Next = l1
            l1 = l1.Next
        } else if l1.Val < l2.Val {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l2.Next
        }
        
        curr = curr.Next
    }
    
    // Advance past node we inserted at beginning
    return head.Next
}
