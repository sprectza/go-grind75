// Problem desc: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
        return nil 
    }

    if root.Val < p.Val && root.Val < q.Val {
        return lowestCommonAncestor(root.Right, p, q) 
    } else if root.Val > p.Val && root.Val > q.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }

    return root 
}
