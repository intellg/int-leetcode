package june_leetcoding_challenge

import "testing"

func TestInvertTree1(t *testing.T) {
	nodes := make([]*TreeNode, 16)
	for i := 0; i < 16; i++ {
		nodes[i] = &TreeNode{
			Val: i,
		}
	}

	root := nodes[10]
	nodes[10].Left = nodes[6]
	nodes[10].Right = nodes[14]
	nodes[6].Left = nodes[3]
	nodes[6].Right = nodes[8]
	nodes[14].Left = nodes[12]
	nodes[14].Right = nodes[15]
	nodes[3].Left = nodes[1]
	nodes[3].Right = nodes[4]
	nodes[8].Left = nodes[7]
	nodes[8].Right = nodes[9]
	nodes[12].Left = nodes[11]
	nodes[12].Right = nodes[13]
	nodes[1].Left = nodes[0]
	nodes[1].Right = nodes[2]
	length := 15
	beforeBase := []int{10, 6, 14, 3, 8, 12, 15, 1, 4, 7, 9, 11, 13, 0, 2}
	afterBase := []int{10, 14, 6, 15, 12, 8, 3, 13, 11, 9, 7, 4, 1, 2, 0}
	doTest(t, root, length, beforeBase, afterBase)
}

func TestInvertTree2(t *testing.T) {
	nodes := make([]*TreeNode, 16)
	for i := 0; i < 16; i++ {
		nodes[i] = &TreeNode{
			Val: i,
		}
	}
	root := nodes[4]
	nodes[4].Left = nodes[2]
	nodes[4].Right = nodes[7]
	nodes[2].Left = nodes[1]
	nodes[2].Right = nodes[3]
	nodes[7].Left = nodes[6]
	nodes[7].Right = nodes[9]
	length := 7
	beforeBase := []int{4, 2, 7, 1, 3, 6, 9}
	afterBase := []int{4, 7, 2, 9, 6, 3, 1}
	doTest(t, root, length, beforeBase, afterBase)
}

func doTest(t *testing.T, root *TreeNode, length int, beforeBase, afterBase []int) {
	arrRoot := make([]*TreeNode, 1)
	arrRoot[0] = root
	beforeInvert := visitWidth(arrRoot)

	invertTree(root)
	afterInvert := visitWidth(arrRoot)

	t.Log(beforeInvert)
	t.Log(afterInvert)

	if length == len(beforeBase) && length == len(afterInvert) {
		for i := 0; i < length; i++ {
			if beforeBase[i] != beforeInvert[i] || afterBase[i] != afterInvert[i] {
				t.Error("Wrong")
				return
			}
		}
	} else {
		t.Error("Wrong")
		return
	}
	t.Log("Correct")
}

func visitWidth(nodes []*TreeNode) (result []int) {
	nextNodes := make([]*TreeNode, 0)
	for _, node := range nodes {
		result = append(result, node.Val)
		if node.Left != nil {
			nextNodes = append(nextNodes, node.Left)
		}

		if node.Right != nil {
			nextNodes = append(nextNodes, node.Right)
		}
	}
	if len(nextNodes) != 0 {
		result = append(result, visitWidth(nextNodes)...)
	}
	return
}
