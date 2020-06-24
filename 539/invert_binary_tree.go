// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/539/week-1-june-1st-june-7th/3347/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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

func main() {
	nodes := make([]*TreeNode, 16)
	for i := 0; i < 16; i++ {
		nodes[i] = &TreeNode{
			Val: i,
		}
	}

	//nodes[10].Left = nodes[6]
	//nodes[10].Right = nodes[14]
	//nodes[6].Left = nodes[3]
	//nodes[6].Right = nodes[8]
	//nodes[14].Left = nodes[12]
	//nodes[14].Right = nodes[15]
	//nodes[3].Left = nodes[1]
	//nodes[3].Right = nodes[4]
	//nodes[8].Left = nodes[7]
	//nodes[8].Right = nodes[9]
	//nodes[12].Left = nodes[11]
	//nodes[12].Right = nodes[13]
	//nodes[1].Left = nodes[0]
	//nodes[1].Right = nodes[2]

	root := nodes[4]
	nodes[4].Left = nodes[2]
	nodes[4].Right = nodes[7]
	nodes[2].Left = nodes[1]
	nodes[2].Right = nodes[3]
	nodes[7].Left = nodes[6]
	nodes[7].Right = nodes[9]

	OutputJson("original.json", root)

	invertTree(root)
	OutputJson("inverted.json", root)
}

func OutputJson(fileName string, root *TreeNode) {
	jsonContent, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = ioutil.WriteFile(fileName, jsonContent, 0644)
}
