package base

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// @description: 将数组转成树状结构, 数组下标i，左节点2i+1, 右节点2i+2,一次遍历完成节点之间的关系，数值为0表示空节点
// @param {[]int} arr
// @return {*}
func SortIntsToTree(arr []int) *TreeNode {

	length := len(arr)

	if length < 0 {
		return nil
	}

	// 节点数组
	nodes := make([]*TreeNode, length)

	for i := length - 1; i >= 0; i-- {
		// 大于0就表示有节点，小于等于0就表示空节点
		if arr[i] > 0 {
			node := &TreeNode{
				Value: arr[i],
			}

			leftIndex := 2*i + 1
			rightIndex := 2*i + 2

			if leftIndex < length {
				node.Left = nodes[leftIndex]
			}

			if rightIndex < length {
				node.Right = nodes[rightIndex]
			}

			nodes[i] = node
		} else {
			nodes[i] = nil
		}
	}

	return nodes[0]
}
