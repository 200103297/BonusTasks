package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	// Value - znachenie
	// Next - address sledueshego noda
}
type Stack struct {
	Peak *Node
	Size int
	// Peak - samaya verhushka stacka
	// Size - velichina stacka
}

func (s *Stack) Pop() int {
	// Checking if there is nothing in stack
	if s.Peak == nil {
		return 0
	}
	// Giving Peak node's value to 'ans'
	// Then replace Peak node with next node
	// Decreasing size by 1
	ans := s.Peak.Value
	s.Peak = s.Peak.Next
	s.Size--
	return ans
}
func (s *Stack) Peek() int {
	// Checking if there is nothing in stack
	if s.Peak == nil {
		return 0
	}
	// Returning Peak's value
	return s.Peak.Value
}
func (s *Stack) Push(value int) {
	// Creating a new node newN with value that we gave in main
	// Then it became a new Peak node because of LIFO
	// Giving the address of previous node to new Peak node
	// Incrementing size of Stack
	newN := &Node{Value: value}
	newN.Next = s.Peak
	s.Peak = newN
	s.Size++
}
func (s *Stack) Clear() {
	// We need to clear only Peak node of stack
	// because of next node's address is in Peak node
	//so if there is no Peak node we can't access to another part of stack
	s.Peak = nil
	s.Size = 0
	fmt.Println("Stack is clear")
}
func (s *Stack) Contains(value int) bool {
	// We create temporary node for checking
	checkingN := s.Peak
	for checkingN != nil {
		// Go through all nodes to find the value
		if checkingN.Value == value {
			return true
		}
		// if we didn't find the value we go to the next node
		checkingN = checkingN.Next
		// This loop will go from top to bottom until it find or not
	}
	return false
}
func (s *Stack) Increment() {
	// We create temporary node
	incrementingN := s.Peak
	for incrementingN != nil {
		// We go through all nodes and just increment it by 1
		incrementingN.Value++
		// Then we go to the next node until we didn't reach last one
		incrementingN = incrementingN.Next
	}
}
func (s *Stack) Print() {
	// We create temporary node
	tempN := s.Peak
	for tempN != nil {
		// We go through all nodes to print their values
		fmt.Print(tempN.Value, " ")
		tempN = tempN.Next
	}
	fmt.Println()
}
func (s *Stack) PrintReverseStack() {
	// In this func we start reversing by recursion
	// we first throw the Peak node to the printReverseNode
	s.printReverseNode(s.Peak)
	fmt.Println()
}
func (s *Stack) printReverseNode(node *Node) {
	// There we check this node which we accept
	// if its equal to null we go to the next node
	// until we get to the last one
	if node == nil {
		return
	}
	s.printReverseNode(node.Next)
	fmt.Print(node.Value, " ")
}
func main() {
	st := &Stack{}
	st.Push(1111)
	st.Push(2222)
	st.Push(3333)
	st.Push(4444)
	st.Push(5555)
	st.Push(6666)
	st.Push(7777)
	fmt.Println("Stack's elements:")
	st.Print()
	fmt.Println("Does stack contain 2222?: ", st.Contains(2222))
	popping := st.Pop()
	fmt.Println("Popping value:", popping)
	fmt.Println("Stack after popping:")
	st.Print()
	peeking := st.Peek()
	fmt.Println("Value after peeking:", peeking)
	st.Increment()
	fmt.Println("Incremented elements of stack:")
	st.Print()
	fmt.Println("Reverse order of stack:")
	st.PrintReverseStack()
	st.Clear()
	st.Print()
}
