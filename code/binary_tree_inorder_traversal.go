package code

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func InOrderTraversal(root *TreeNode) []int {
	var nums = []int{}
	if root == nil {
		return nums
	}

	nums = append(nums, InOrderTraversal(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, InOrderTraversal(root.Right)...)
	return nums
}
