package main

import (
	"fmt"
)

//Singly Linked List

type NodeSingly struct{
	Info interface{}
	Next *NodeSingly
}


type List struct{
	Head *NodeSingly
}

func NewSinglyNode(info interface{})*NodeSingly{
	return &NodeSingly{Info: info,Next: nil}
}

func(list *List) InsertAtTheEnd(d interface{}){
	node := NewSinglyNode(d)

	if list.Head ==nil{
		list.Head = node

	}else{
		current := list.Head

		for current.Next !=nil{
			current = current.Next
		}
		current.Next = node
	}
}

func(list *List) InsertAtTheBegining(d interface{}){
	node := NewSinglyNode(d)

	node.Next = list.Head
	list.Head = node
}


func (list *List) Show(){
	current:=list.Head

	for current.Next !=nil{
		fmt.Printf("-> %v", current.Info)
		current = current.Next
	}
}

func(list *List) DeleteInfo(d interface{}){
	if list.Head == nil {
		return
	}
	
if list.Head.Info == d {
		list.Head = list.Head.Next
		return
	}
	prev := list.Head
	current := list.Head.Next
	for current != nil {
		if current.Info == d {
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}
}