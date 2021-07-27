package main

import "reflect"

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

/*

package main

import "reflect"

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}
func (n Node) IsStructureEmpty() bool {
	return reflect.DeepEqual(n, Node{})
}
func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array,n.Name)
	temp := *n
	for i := 0; i < len(temp.Children);i++{
		if temp.Children[i].IsStructureEmpty(){
		}else{
			array = n.Children[i].DepthFirstSearch(array)
		}
	}
	return array
}


*/
func (n Node) IsStructureEmpty() bool {
	return reflect.DeepEqual(n, Node{})
}
func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}
	return array
}
