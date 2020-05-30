package s

var memo = make([][]*TreeNode, 20)

func allPossibleFBT(N int) []*TreeNode {
	if N == 1 {
		return []*TreeNode{&TreeNode{0, nil, nil}}
	}
	var res []*TreeNode
	for i := 1; i <= N-1; i = i + 2 {
		if memo[i] == nil {
			memo[i] = allPossibleFBT(i)
		}
		if memo[N-i-1] == nil {
			memo[N-i-1] = allPossibleFBT(N - i - 1)
		}
		for _, l := range memo[i] {
			for _, r := range memo[N-i-1] {
				res = append(res, &TreeNode{0, l, r})
			}
		}
	}
	return res
}
