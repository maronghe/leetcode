package s

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	var q []*TreeNode
	if root != nil {
		q = append(q, root) // add root
	}

	for len(q) > 0 {
		n := len(q)              // get size of level
		var level []int          // level result
		for k := 0; k < n; k++ { // execute level's size times
			// pop()
			node := q[0]
			q = q[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}

	return res
}
