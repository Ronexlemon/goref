package main

/***
-> Whats is Heap? -> A special binary tree in which all its levels are completely filled up except possibly the last one and all nodes are filled from left to right
types of binary tree:
1. Full Binary Tree -> every node has two children
2. Complete Binary Tree -> Only the last level is not filled up


HEAP PROPERTY:
   1. Min heap: the value of child nodes is greater than or equal to the value of its parent.The root node has the minimum value
   2. Max Heap: The value of child nodes is smaller than or equal to the value of its parent.The root node has the maximum value

***/
/*****Min HEAP**********/
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) swap(i, j int)       { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}){
	*h = append(*h, x.(int))

}

func (h *MinHeap) Pop()interface{}{
	old := *h

	n:= len(old)

	x:= old[n-1] //the one i want to pop
	*h=old[:n-1]

	return x
}

func MinnHeap(){


}

/****
Sift Down:-> Move the value down the tree by successively exchanging the value with its smaller(for min heap)/larger(for max heap) childnode
Sift up:-> move the value up the tree by successively exchanging the value with its parent node.

****/

//heapyfy an array

func (h *MinHeap) BuildHeap(array []int) {
	*h = array
	lastNonLeaf := (len(array) - 2) / 2
	for i := lastNonLeaf; i >= 0; i-- {
		h.siftDown(i, len(array)-1)
	}
}


func (h *MinHeap) siftDown(currentIdx int, endIdx int){

	leftChildIdx := currentIdx*2+1

	for leftChildIdx <= endIdx{
		rightChildIdx := currentIdx*2+2

		if rightChildIdx > endIdx{
			rightChildIdx =-1
		}
		//get the smaller child node to swap
         idxtoSwap := leftChildIdx
		if rightChildIdx !=-1 && (*h)[rightChildIdx] < (*h)[leftChildIdx]{
            idxtoSwap = rightChildIdx
		}

		//check if value of swap node is less than the value at currentIdx
		if(*h)[idxtoSwap] < (*h)[currentIdx]{
			h.swap(idxtoSwap,currentIdx)
			currentIdx = idxtoSwap

			leftChildIdx = currentIdx*2+1
		}else{
			return
		}
	}
}

func(h *MinHeap) siftUp(){
	currentIdx := len(*h)-1

	parentIdx := (currentIdx-1)/2

	for currentIdx > 0 && (*h)[currentIdx] < (*h)[parentIdx]{
		h.swap(currentIdx,parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx-1)/2

	}
}

//insert a new value to the end of the tree and update heap ordering

func (h *MinHeap) Insert(value int){
	*h = append(*h, value)
	h.siftUp()
}

//remove and return the minimum value and update heap ordering

func (h *MinHeap) Remove()int{
	n :=len(*h)

	//swap the first element and the last element in the array

	h.swap(0,n-1)
	valueToRemove := (*h)[n-1]

	//pop the last element in the array
	*h = (*h)[:n-1]

	//call siftdown to update heap ordering
	h.siftDown(0,n-2)
	return valueToRemove
}