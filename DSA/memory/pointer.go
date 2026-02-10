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
}


func UpdateName(name *string){
	*name = "Yes Baanah"
}