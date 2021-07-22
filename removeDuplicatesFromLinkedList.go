package main

// LinkedList This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	temp := linkedList
	for temp.Next != nil {
		temp2 := temp.Next
		if temp.Value == temp2.Value {
			temp.Next = temp2.Next
		} else {
			temp = temp.Next
		}
		if temp.Next == nil {
			break
		}
	}

	return linkedList
}
