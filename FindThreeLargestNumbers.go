package main

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Push(str int) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}
func FindThreeLargestNumbers(array []int) []int {
	myStack := Stack{}
	myStack2 := Stack{}
	for i := 0; i <= len(array)-1; i++ {
		if myStack.IsEmpty() {
			myStack.Push(array[i])
		} else {
			num := myStack.Pop()
			if array[i] >= num {
				myStack.Push(num)
				myStack.Push(array[i])
			} else {
				for array[i] < num {
					myStack2.Push(num)
					if myStack.IsEmpty() {
						break
					}
					num = myStack.Pop()
				}
				if num < array[i] {
					myStack.Push(num)
				}
				myStack.Push(array[i])
				for !myStack2.IsEmpty() {
					myStack.Push(myStack2.Pop())
				}
			}
		}
	}
	mySlice := []int{}
	for x := 2; x >= 0; x-- {
		mySlice = append([]int{myStack.Pop()}, mySlice...)
	}
	return mySlice
}
