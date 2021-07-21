package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	c:= 0
	 Walk(tree,target,&c)
	 return c
}

func Walk(t *BST, target int,c  *int)  {

	if t == nil {
		return
	}



	// walk the left side first
	Walk(t.Left, target,c)

	if *c == 0 {
		*c = t.Value
	}else{
		if math.Abs(float64(target-t.Value)) < math.Abs(float64(target-*c)){
			*c = t.Value
		}
	}

	// walk the right side last
	Walk(t.Right, target,c)







}

