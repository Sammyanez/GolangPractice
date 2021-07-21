package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	sumDepth := 0
	depthCalulator(root,&sumDepth,0)
	return sumDepth
}

func depthCalulator(root *BinaryTree, depth *int,num int) {
	if root == nil{
		return
	}

	if num == 0{num++}else{
	*depth += num
	num++}
	if root.Left == nil && root.Right ==nil{
	return}


	depthCalulator(root.Left , depth ,num )
	depthCalulator(root.Right , depth ,num )
	num = 0


}
