package main

import "fmt"

/***
declaration
type and struct key word
**/

type Person struct{
	Name string
	Age int
	Job string
	Salary int
}

func logPerson(pern Person){
	fmt.Println("The Name:", pern.Name)
	fmt.Println("The Age:", pern.Age)
	fmt.Println("The Job:", pern.Job)
	fmt.Println("The Sal:", pern.Salary)
}

func InitPerson(){
	var person Person
	person2:=Person{Name: "lemn",Age: 25,Job: "Yoolw",Salary: 100000000}
	person.Age =24
	person.Name ="lemonr"
	person.Job ="block infra"
	person.Salary= 200000
	logPerson(person)
	logPerson(person2)
}