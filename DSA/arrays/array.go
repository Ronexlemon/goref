package main

import "fmt"

func Arrays(){
	a:=[10]int{1,2,3,4,5,6,7,8,9,10}
	var b [10]string

	a[0]=0
	b[1]="Hello"
	b[5]="you"
	fmt.Println("Array a:",a)
	fmt.Println("Array b:",b)
}

//Declare an array
//1. With the var keyword
  //var array_name [length] datatype
var newArray [10]int

//2. var array_name = [length]datatype{values} * here length is defined
//3. var array_name =[...]datatype{values} * here length is infered
//4. array_name:=[length]datatype{values} * the length is defined
//5. array_name:=[...]datatype{values} * the length is infered


func BaseOnDeclaration(){
	var k [10]int
	var a =[5]int{1,2,3,4,5}
	var aa =[...]string{"hellow","yes","yooh","whi"}

	b:=[5]int{1,2,3,4,5}
	bb:=[...]string{"yes","yooh","yes","yooh"}
	for key,_:=range k{
		k[key]+=1;
	}
	fmt.Println("The A:",a)
	fmt.Println("The AA",aa)
	fmt.Println("The B:",b)
	fmt.Println("The BB:",bb)
	fmt.Println("The K:",k)
}

func InitSpecificElements(){
	arr1:=[5]int{1:10,3:30}
	fmt.Println("The Init Arr:",arr1)
}

