package main

import "fmt"

/*
Features of Linked List
1. Memory Allocation:
   Nodes are stored non-contiguously in memory and linked via pointers.

2. Access Pattern:
   Sequential access only — traversal from the head is required.

3. Insertion & Deletion:
   O(1) at the beginning (and end if tail is tracked), O(n) elsewhere.

4. Dynamic Size:
   Grows and shrinks dynamically but has extra memory overhead for pointers.

5. Cache Performance:
   Poor cache locality due to scattered memory allocation.
*/

// Node represents a single element in the list
type Node struct {
	Data int
	Next *Node
}

// SinglyLinkedList represents the list
type SinglyLinkedList struct {
	Head *Node
	Tail *Node
}

// NewNode creates a new node
func NewNode(data int) *Node {
	return &Node{Data: data}
}

// Append adds a node to the end of the list (O(1))
func (list *SinglyLinkedList) Append(data int) {
	node := NewNode(data)

	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return
	}

	list.Tail.Next = node
	list.Tail = node
}

// Remove deletes the first occurrence of data
func (list *SinglyLinkedList) Remove(data int) {
	if list.Head == nil {
		return
	}

	// If head needs to be removed
	if list.Head.Data == data {
		list.Head = list.Head.Next
		if list.Head == nil {
			list.Tail = nil
		}
		return
	}

	current := list.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			if current.Next == nil {
				list.Tail = current
			}
			return
		}
		current = current.Next
	}
}

// Display prints all nodes in the list
func (list *SinglyLinkedList) Display() {
	current := list.Head

	for current != nil {
		fmt.Printf("Node: %d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func CallSinglyLinkedList() {
	list := SinglyLinkedList{}

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)

	fmt.Println("Initial list:")
	list.Display()

	list.Remove(20)
	fmt.Println("After removing 20:")
	list.Display()

	list.Remove(10)
	fmt.Println("After removing head (10):")
	list.Display()

	list.Remove(40)
	fmt.Println("After removing tail (40):")
	list.Display()
}


//Doubly Linked List(DLS)

/****
This are enhancements over the SLLs for which additional pointer (prev) is introduced to point the previous node, thereby setting a bi-directional structure which gives a lot of advantages over SLLs. Insertion and deletion at both ends of the DLLs occur in constant time


***/

type DoublyNode struct{
	Data int
	Prev *DoublyNode
	Next *DoublyNode
}

func NewDoublyNode(data int)*DoublyNode{
	return &DoublyNode{Data: data,Prev: nil,Next: nil}

}

type DoublyLinkedList struct{
	Head *DoublyNode
	Tail *DoublyNode
}

//Insertion at the end

func(list *DoublyLinkedList) InsertAtEnd(data int){
	node := NewDoublyNode(data)

	// Empty list
	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return
	}

	list.Tail.Next = node
	node.Prev = list.Tail
	list.Tail = node
}