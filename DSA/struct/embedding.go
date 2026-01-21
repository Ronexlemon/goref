package main

import "fmt"

type base struct{
	num int
}

type container struct{
	base
	str string
}

func(b base) describe()string{
	return fmt.Sprintf("base with num=%v",b.num)
}

func InitBase(){
	b:=base{}
	fmt.Println(b.describe())
}

func Embedding(){
	co:= container{
		base:base{
			num: 20,
		},
		str: "Some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n",co.num,co.str)
	fmt.Println("also num:",co.base.num)
	//since container embeds base, the methods of base also become methods of a container
	fmt.Println(co.describe())


	type describe interface{
		describe() string
	}
	var d describe = co
	fmt.Println(d.describe())
}