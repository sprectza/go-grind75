// Problem desc: https://leetcode.com/problems/merge-two-sorted-lists/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy 

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            current.Next = l1 
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }

        current = current.Next 
    }

    if l1 != nil {
        current.Next = l1
    } else {
        current.Next = l2
    }

    return dummy.Next 
}
