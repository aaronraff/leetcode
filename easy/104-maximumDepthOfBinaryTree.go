/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil { return 0 }
    
    leftDepth := float64(1 + maxDepth(root.Left))
    rightDepth := float64(1 + maxDepth(root.Right))
    
    return int(math.Max(leftDepth, rightDepth))
}
