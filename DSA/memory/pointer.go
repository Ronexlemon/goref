package main

import "fmt"



var Num *int

type Person struct{
	Name string
}

func InitPerson()*Person{
	return &Person{}
}

func Assign(val int){
	Num = &val
}


func Get()*int{
	return Num
}

func Pointer(){
	Assign(10)
	fmt.Println("The value is",*Get())
	name:= "Lemonr"
	UpdateName(&name)
	fmt.Println("The updated Name is",name)
	p := InitPerson()

	p.Name = "John"
	println("The Struct person",p.Name)
	CheckSwap()
}


func UpdateName(name *string){
	*name = "Yes Baanah"
}


func Swap(a,b *int){
	placeHolder := *a

	*a = *b 

	*b = placeHolder


}
func SwapB (a,b *int){
	*a,*b =*b,*a
}

func CheckSwap(){
	num1 := 10
	num2:=20
	//Swap(&num1,&num2)
	SwapB(&num1,&num2)
	fmt.Printf("The value of num1 :%d and num2 : %d",num1,num2)

}