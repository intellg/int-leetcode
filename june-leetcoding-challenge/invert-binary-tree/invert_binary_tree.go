// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/539/week-1-june-1st-june-7th/3347/
package invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left != nil {
		invertTree(node.Left)
	}
	if node.Right != nil {
		invertTree(node.Right)
	}

	p := node.Left
	node.Left = node.Right
	node.Right = p
	return node
}
