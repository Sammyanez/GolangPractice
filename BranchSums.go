package main
/*
BinaryTree This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	var sums []int
	calculateBranchSums(root,0,&sums )

	return sums
}

func calculateBranchSums(root *BinaryTree, runningSum int, sums *[]int) {
	if root == nil{
		return
	}

	runningSum += root.Value
	if root.Left == nil && root.Right == nil{
		*sums = append(*sums,runningSum)
		return
	}

	calculateBranchSums(root.Left,runningSum,sums)
	calculateBranchSums(root.Right,runningSum,sums)

}
*/