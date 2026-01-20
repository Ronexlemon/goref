package main

import "fmt"


func Slice(){
	/***
	Syntax
	slice_name:=[]datatype{values}
	
	****/
	slice1:=[]int{}
	fmt.Println("The length:", len(slice1))
	fmt.Println("The Capacity:",cap(slice1))

	fmt.Println("Content:",slice1)

	slice2 :=[]string{"Yello","blu","blac"}
	fmt.Println("The Slice2")
	fmt.Println("The length:", len(slice2))
	fmt.Println("The Capacity:",cap(slice2))

	fmt.Println("Content:",slice2)
	slice2 = append(slice2, "Green")
	fmt.Println("After appending:")
	fmt.Println("Content:",slice2)
}

//creating a slice with make

func SliceWithMake(){
	slice1:=make([]int,21)

	fmt.Println("The Slice:",slice1)
	fmt.Println("The Length:", len(slice1))
	fmt.Println("The capacity:",cap(slice1))
	slice1[20]=10
	fmt.Println("Content:",slice1)
}

//Append two slices

func AppendTwoSlices(){
	slice1:=[]int{1,2,3,4,5,6}
	slice2:=[]int{7,8,9,10}
	slice3:= append(slice1,slice2...)
	fmt.Println("Slice 3",slice3)
}

/***
The copy() function creates a new underlying array with only the required elements for the slice.
this will reduce the memory used for the program
copy(dest,src)
***/

func Slicecopy(){
	numbers:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

	fmt.Printf("numbers = %v\n",numbers)
	fmt.Printf("length = %d\n",len(numbers))
	fmt.Printf("capacity = %d\n",cap(numbers))

	//create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy :=make([]int,len(neededNumbers))

	copy(numbersCopy,neededNumbers)

	fmt.Printf("numbers = %v\n",neededNumbers)
	fmt.Printf("length = %d\n",len(neededNumbers))
	fmt.Printf("capacity = %d\n",cap(neededNumbers))


}
