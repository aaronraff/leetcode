/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    dummyL := &ListNode{}
    dummyR := &ListNode{}
    currL, currR := dummyL, dummyR
    curr := head
    
    for curr != nil {
        if curr.Val < x {
            currL.Next = curr
            currL = currL.Next
        } else {
            currR.Next = curr
            currR = currR.Next
        }
        
        // Advance
        curr = curr.Next
    }
    
    // Terminate list
    currR.Next = nil
    
    // Join the two partitions
    currL.Next = dummyR.Next
    return dummyL.Next
}
