package collections

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeToStr(root *TreeNode) string {
	arr := TreeToNodeArr(root)
	for len(arr) > 0 && arr[len(arr)-1] == nil {
		arr = arr[:len(arr)-1]
	}
	result := make([]string, 0, len(arr))
	for _, node := range arr {
		result = append(result, ValToStr(node))
	}
	return strings.Join(result, ",")
}

func StrToTree(s string) *TreeNode {
	if s == "" {
		return nil
	}
	return StrArrToTree(strings.Split(s, ","), 0)
}

func TreeToNodeArr(root *TreeNode) []*TreeNode {
	result := make([]*TreeNode, 0)
	result = append(result, root)
	offset := 0
	for offset < len(result) {
		count := len(result) - offset
		for i := 0; i < count; i++ {
			if result[offset+i] != nil {
				result = append(result, result[offset+i].Left, result[offset+i].Right)
			}
		}
		offset += count
	}
	return result
}

func ValToStr(node *TreeNode) string {
	if node == nil {
		return "null"
	}
	return strconv.Itoa(node.Val)
}

func StrArrToTree(values []string, i int) *TreeNode {
	if i >= len(values) || values[i] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(values[i])
	node := &TreeNode{Val: val}
	node.Left = StrArrToTree(values, 2*i+1)
	node.Right = StrArrToTree(values, 2*i+2)
	return node
}
