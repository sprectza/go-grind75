/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return height(root) != -1
}

func height(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftH := height(root.Left) 
    rightH := height(root.Right) 

    if leftH == -1 || rightH == -1 || math.Abs(float64(leftH - rightH)) > 1 {
        return -1
    }

    return 1 + max(leftH, rightH)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
