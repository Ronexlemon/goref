package main

import "fmt"

type Clans interface{
	Speak(string) string
	Skills(interface{}) interface{}

}

type Mountainers struct{}

func(m *Mountainers) Speak(language string)(string){
	return language

}

func (m *Mountainers) Skills(skills interface{})(interface{}){
	return skills
}

func InterfaceAndMethods(){
	var clans Clans
	var skills =[]string{"hunting","gathering","cooking"}
	clans = &Mountainers{}
	fmt.Println(clans.Skills(skills))
	fmt.Println(clans.Speak("Yoruba"))
}