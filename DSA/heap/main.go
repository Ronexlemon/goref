package main

import "fmt"

func main(){
	fmt.Println("Heap........")
	array :=[]int{9,32,45,3,94,8,60,104,56,4,58,98}
	minHeap := &MinHeap{}
	minHeap.BuildHeap(array)
	//apply  remove method
	valueToRemove := minHeap.Remove()
	fmt.Println("The root value",valueToRemove)
	fmt.Println("min Heap after Remove root",*minHeap)

	minHeap.Insert(50)
	fmt.Println("min Heap after insert ",*minHeap)

}
